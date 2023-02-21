/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package manageiq

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/PDeXchange/pac/internal/pkg/client/iam"
	"github.com/PDeXchange/pac/internal/pkg/client/powervs"
	"github.com/PDeXchange/pac/internal/pkg/client/vpc"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	manageiqv1alpha1 "github.com/PDeXchange/pac/apis/manageiq/v1alpha1"
)

var (
	ErrorPoolNotFound     = fmt.Errorf("pool not found")
	ErrorListenerNotFound = fmt.Errorf("listener not found")
)

// ServiceReconciler reconciles a Service object
type ServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=manageiq.pac.io,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=manageiq.pac.io,resources=services/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=manageiq.pac.io,resources=services/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Service object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	l.Info("Hello")

	service := &manageiqv1alpha1.Service{}

	if err := r.Get(ctx, req.NamespacedName, service); err != nil {
		l.Error(err, "unable to fetch CronJob")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	switch service.Spec.Type {
	case manageiqv1alpha1.ServiceTypeVM:
		auth, err := iam.GetIAMAuth()
		if err != nil {
			l.Error(err, "failed to authenticate")
			return ctrl.Result{}, err
		}
		if err := r.reconcileVM(ctx, auth, service); err != nil {
			l.Error(err, "unable to reconcileVM")
			return ctrl.Result{}, err
		}
	default:
		l.Error(fmt.Errorf("unknown service type"), string(service.Spec.Type))
		return ctrl.Result{}, fmt.Errorf("unknown service type")
	}

	l.Info("service object", "service", service.Spec)

	return ctrl.Result{}, nil
}

func (r *ServiceReconciler) reconcileVM(ctx context.Context, auth core.Authenticator, service *manageiqv1alpha1.Service) error {
	l := log.FromContext(ctx)
	l.V(4).Info("in the reconcileVM")

	defer func() {
		if err := r.Status().Update(ctx, service); err != nil {
			l.Error(err, "unable to update Service status")
		}
	}()

	// This will initialise the status with ports from spec
	if err := reconcilePorts(ctx, r, service); err != nil {
		return err
	}

	// This will update the status with the right mac address
	if err := reconcileNetwork(ctx, r, service); err != nil {
		return err
	}

	// reconcileIPAddress reconcile the IP addresses from the dhcp server's leases and assign to status
	if err := reconcileIPAddress(ctx, r, service); err != nil {
		return err
	}

	service.Status.SetReady()

	if err := reconcileIngress(ctx, service); err != nil {
		return err
	}

	return nil
}

func reconcileIngress(ctx context.Context, service *manageiqv1alpha1.Service) error {
	l := log.FromContext(ctx)

	l.V(4).Info("in the reconcileIngress")

	client, err := vpc.NewClient(ctx, vpc.Options{Region: service.Spec.VirtualMachine.VPC.Region})
	if err != nil {
		return err
	}

	lbID := service.Spec.VirtualMachine.VPC.Loadbalancer

	// Step1: Ensure to create a pool
	findPool := func(port uint) (*vpcv1.LoadBalancerPool, error) {
		pools, _, err := client.ListLoadBalancerPools(&vpcv1.ListLoadBalancerPoolsOptions{LoadBalancerID: &lbID})
		if err != nil {
			return nil, err
		}
		for _, pool := range pools.Pools {
			if fmt.Sprintf("%s-%d", service.Spec.VirtualMachine.Name, port) == *pool.Name {
				return &pool, nil
			}
		}
		return nil, fmt.Errorf("pool not found")
	}
	for _, net := range service.Status.VirtualMachine.Networks {
		if _, err := findPool(net.Port); err != nil && err != ErrorPoolNotFound {
			continue
		} else if err == ErrorPoolNotFound {
			l.Info("need to write this code to create pool")
		}
	}

	// Step2: Ensure to create a listener
	findListener := func(name string) (*vpcv1.LoadBalancerListener, error) {
		l, _, err := client.ListLoadBalancerListeners(&vpcv1.ListLoadBalancerListenersOptions{LoadBalancerID: &lbID})
		if err != nil {
			return nil, err
		}

		for _, listener := range l.Listeners {
			if *listener.DefaultPool.Name == name {
				return &listener, nil
			}
		}
		return nil, ErrorListenerNotFound
	}
	for i, net := range service.Status.VirtualMachine.Networks {
		listener, err := findListener(fmt.Sprintf("%s-%d", service.Spec.VirtualMachine.Name, net.Port))
		if err != nil && err != ErrorListenerNotFound {
			continue
		} else if err == ErrorListenerNotFound {
			pool, err := findPool(net.Port)
			if err != nil {
				continue
			}

			createOpt := &vpcv1.CreateLoadBalancerListenerOptions{
				LoadBalancerID:  core.StringPtr(lbID),
				ConnectionLimit: core.Int64Ptr(15000),
				Port:            core.Int64Ptr(int64(40002)),
				Protocol:        core.StringPtr("tcp"),
				DefaultPool:     &vpcv1.LoadBalancerPoolIdentity{ID: pool.ID},
			}
			listener, _, err = client.CreateLoadBalancerListener(createOpt)
			if err != nil {
				l.Error(err, "failed to CreateLoadBalancerListener")
				continue
			}
		}
		if listener != nil && listener.Port != nil {
			service.Status.VirtualMachine.Networks[i].TargetPort = uint(*listener.Port)
		}
	}

	return nil
}

func reconcileIPAddress(ctx context.Context, r *ServiceReconciler, service *manageiqv1alpha1.Service) error {
	l := log.FromContext(ctx)

	l.V(4).Info("in the reconcileIPAddress")

	vmStatus := &service.Status.VirtualMachine

	if vmStatus.NetworkName == "" {
		l.Error(fmt.Errorf("network information is not yet available"), "retry after sometime")
		return fmt.Errorf("network information is not yet available")
	}
	client, err := powervs.NewClient(ctx, powervs.Options{
		CloudInstanceID: service.Spec.VirtualMachine.CloudInstanceID,
		Zone:            service.Spec.VirtualMachine.Zone})
	if err != nil {
		return err
	}

	dhcpservers, err := client.GetAllDHCPServers()
	if err != nil {
		return err
	}

	for _, dhcpserver := range dhcpservers {
		if *dhcpserver.Network.Name == vmStatus.NetworkName {
			s, err := client.GetDHCPServer(*dhcpserver.ID)
			if err != nil {
				return err
			}
			for _, lease := range s.Leases {
				if *lease.InstanceMacAddress == vmStatus.MACAddress && *lease.InstanceIP != "" {
					vmStatus.IPAddress = *lease.InstanceIP
					return nil
				}
			}
		}
	}

	return fmt.Errorf("no lease found for the assigned mac address")
}

func reconcileNetwork(ctx context.Context, r *ServiceReconciler, service *manageiqv1alpha1.Service) error {
	l := log.FromContext(ctx)
	l.V(4).Info("in the reconcileNetwork")
	client, err := powervs.NewClient(ctx, powervs.Options{
		CloudInstanceID: service.Spec.VirtualMachine.CloudInstanceID,
		Zone:            service.Spec.VirtualMachine.Zone})
	if err != nil {
		return err
	}

	in, err := func() (*models.PVMInstanceReference, error) {
		ins, err := client.GetAllInstance()
		if err != nil {
			return nil, err
		}
		for _, in := range ins.PvmInstances {
			if service.Spec.VirtualMachine.Name == *in.ServerName {
				l.V(3).Info("found the vm!", "instance", in)
				return in, nil
			}
		}
		return nil, fmt.Errorf("vm not found")
	}()

	if err != nil {
		return err
	}

	service.Status.SetVirtualMachineStatusInstanceID(in)
	service.Status.SetVirtualMachineStatusMACAddress(in)

	return nil
}

func reconcilePorts(ctx context.Context, r *ServiceReconciler, service *manageiqv1alpha1.Service) error {
	l := log.FromContext(ctx)
	l.V(4).Info("in the reconcilePorts")

	addPorts := func(port uint) {
		for _, nw := range service.Status.VirtualMachine.Networks {
			if nw.Port == port {
				return
			}
		}
		service.Status.VirtualMachine.Networks = append(service.Status.VirtualMachine.Networks, manageiqv1alpha1.NetworkStatus{Port: port})
	}

	for _, port := range service.Spec.VirtualMachine.Ports {
		addPorts(port)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&manageiqv1alpha1.Service{}).
		Complete(r)
}
