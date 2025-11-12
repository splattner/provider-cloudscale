// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	network "github.com/splattner/provider-cloudscale/internal/controller/cluster/cloudscale/network"
	objectsuser "github.com/splattner/provider-cloudscale/internal/controller/cluster/cloudscale/objectsuser"
	server "github.com/splattner/provider-cloudscale/internal/controller/cluster/cloudscale/server"
	subnet "github.com/splattner/provider-cloudscale/internal/controller/cluster/cloudscale/subnet"
	providerconfig "github.com/splattner/provider-cloudscale/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.Setup,
		objectsuser.Setup,
		server.Setup,
		subnet.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.SetupGated,
		objectsuser.SetupGated,
		server.SetupGated,
		subnet.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
