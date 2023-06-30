/*
Copyright 2023.

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

package controllers

import (
	"context"

	"github.com/PDeXchange/pac/apis/app/v1alpha1"
	"github.com/PDeXchange/pac/controllers/util"
	"github.com/PDeXchange/pac/internal/pkg/client/platform"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/IBM/go-sdk-core/v5/core"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/PDeXchange/pac/internal/pkg/client/powervs"
)

const (
	catalogController = "catalog"
	serviceController = "service"
)

type ControllerScopeParams struct {
	Logger  logr.Logger
	Client  client.Client
	Type    string
	Catalog *v1alpha1.Catalog
	Debug   bool
}

type ControllerScope struct {
	logr.Logger
	Client         client.Client
	Catalog        *v1alpha1.Catalog
	PowerVSClient  *powervs.Client
	PlatformClient *platform.Client
}

type CatalogScopeParams struct {
	ControllerScopeParams
}

type ServiceScopeParams struct {
	ControllerScopeParams
	Service *v1alpha1.Service
}

type CatalogScope struct {
	ControllerScope
	catalogPatchHelper *patch.Helper
}

type ServiceScope struct {
	ControllerScope
	servicePatchHelper *patch.Helper
	Service            *v1alpha1.Service
}

func NewCatalogScope(ctx context.Context, params CatalogScopeParams) (*CatalogScope, error) {
	scope := &CatalogScope{}

	ctrlScope, err := NewControllerScope(ctx, params.ControllerScopeParams)
	if err != nil {
		return scope, errors.Wrap(err, "failed to init controller scope")
	}

	scope.ControllerScope = *ctrlScope

	catalogHelper, err := patch.NewHelper(params.Catalog, params.Client)
	if err != nil {
		return scope, errors.Wrap(err, "failed to init patch helper")
	}
	scope.catalogPatchHelper = catalogHelper

	return scope, nil
}

func NewServiceScope(ctx context.Context, params ServiceScopeParams) (*ServiceScope, error) {
	scope := &ServiceScope{}

	ctrlScope, err := NewControllerScope(ctx, params.ControllerScopeParams)
	if err != nil {
		err = errors.Wrap(err, "failed to init controller scope")
		return scope, err
	}
	scope.ControllerScope = *ctrlScope

	if params.Service == nil {
		err = errors.New("service is required when creating a ServiceScope")
		return scope, err
	}
	scope.Service = params.Service

	serviceHelper, err := patch.NewHelper(params.Service, params.Client)
	if err != nil {
		err = errors.Wrap(err, "failed to init patch helper")
		return scope, err
	}
	scope.servicePatchHelper = serviceHelper

	return scope, nil
}

func NewControllerScope(ctx context.Context, params ControllerScopeParams) (*ControllerScope, error) {
	scope := &ControllerScope{}

	if params.Client == nil {
		return scope, errors.New("client is required when creating a CatalogScope")
	}
	scope.Client = params.Client

	if params.Logger == (logr.Logger{}) {
		params.Logger = zap.New()
	}
	scope.Logger = params.Logger

	if params.Catalog == nil {
		return scope, errors.New("catalog is required when creating a scope for catalog and service controller")
	}
	scope.Catalog = params.Catalog

	platformClient, err := platform.NewClient()
	if err != nil {
		return scope, errors.Wrap(err, "error creating platform services client")
	}
	scope.PlatformClient = platformClient

	var cloudInstanceID, zone, accountID string
	switch params.Catalog.Spec.Type {
	case v1alpha1.CatalogTypeVM:
		cloudInstanceID, zone, accountID, err = util.ParsePowerVSCRN(params.Catalog.Spec.VM.CRN)
		if err != nil {
			return scope, err
		}
	}

	powerVSClient, err := powervs.NewClient(ctx, powervs.Options{
		AccountID:       accountID,
		CloudInstanceID: cloudInstanceID,
		Zone:            zone,
		Debug:           params.Debug})
	if err != nil {
		return scope, errors.Wrap(err, "failed to create powervs client")
	}
	scope.PowerVSClient = powerVSClient

	if params.Debug {
		core.SetLoggingLevel(core.LevelDebug)
	}

	return scope, nil
}

// PatchObject persists the catalog/service configuration and status.
func (m *CatalogScope) PatchCatalogObject() error {
	return m.catalogPatchHelper.Patch(context.TODO(), m.Catalog)
}

// PatchObject persists the catalog/service configuration and status.
func (m *ServiceScope) PatchServiceObject() error {
	return m.servicePatchHelper.Patch(context.TODO(), m.Service)
}
