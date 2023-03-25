package vpc

import (
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

type VPC interface {
	ListLoadBalancers(options *vpcv1.ListLoadBalancersOptions) (*vpcv1.LoadBalancerCollection, *core.DetailedResponse, error)
	GetLoadBalancer(options *vpcv1.GetLoadBalancerOptions) (*vpcv1.LoadBalancer, *core.DetailedResponse, error)
	ListLoadBalancerPools(options *vpcv1.ListLoadBalancerPoolsOptions) (*vpcv1.LoadBalancerPoolCollection, *core.DetailedResponse, error)
	DeleteLoadBalancerPool(options *vpcv1.DeleteLoadBalancerPoolOptions) (*core.DetailedResponse, error)
	ListLoadBalancerListeners(options *vpcv1.ListLoadBalancerListenersOptions) (*vpcv1.LoadBalancerListenerCollection, *core.DetailedResponse, error)
	CreateLoadBalancerListener(options *vpcv1.CreateLoadBalancerListenerOptions) (*vpcv1.LoadBalancerListener, *core.DetailedResponse, error)
	CreateLoadBalancerPoolMember(options *vpcv1.CreateLoadBalancerPoolMemberOptions) (*vpcv1.LoadBalancerPoolMember, *core.DetailedResponse, error)
	DeleteLoadBalancerPoolMember(options *vpcv1.DeleteLoadBalancerPoolMemberOptions) (*core.DetailedResponse, error)
	ListLoadBalancerPoolMembers(options *vpcv1.ListLoadBalancerPoolMembersOptions) (*vpcv1.LoadBalancerPoolMemberCollection, *core.DetailedResponse, error)
	DeleteLoadBalancerListener(options *vpcv1.DeleteLoadBalancerListenerOptions) (*core.DetailedResponse, error)
}
