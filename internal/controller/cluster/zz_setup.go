// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	awsiamrole "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/awsiamrole"
	crowdsettings "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/crowdsettings"
	globalrole "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/globalrole"
	group "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/group"
	httpssosettings "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/httpssosettings"
	oidcconfiguration "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/oidcconfiguration"
	permission "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/permission"
	reverseproxy "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/reverseproxy"
	samlsettings "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/samlsettings"
	scimgroup "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/scimgroup"
	scimuser "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/scimuser"
	workersservice "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/platform/workersservice"
	providerconfig "github.com/hmlkao/provider-jfrog-platform/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awsiamrole.Setup,
		crowdsettings.Setup,
		globalrole.Setup,
		group.Setup,
		httpssosettings.Setup,
		oidcconfiguration.Setup,
		permission.Setup,
		reverseproxy.Setup,
		samlsettings.Setup,
		scimgroup.Setup,
		scimuser.Setup,
		workersservice.Setup,
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
		awsiamrole.SetupGated,
		crowdsettings.SetupGated,
		globalrole.SetupGated,
		group.SetupGated,
		httpssosettings.SetupGated,
		oidcconfiguration.SetupGated,
		permission.SetupGated,
		reverseproxy.SetupGated,
		samlsettings.SetupGated,
		scimgroup.SetupGated,
		scimuser.SetupGated,
		workersservice.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
