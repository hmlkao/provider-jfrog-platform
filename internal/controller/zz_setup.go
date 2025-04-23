// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	globalrole "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/globalrole"
	group "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/group"
	httpssosettings "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/httpssosettings"
	oidcconfiguration "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/oidcconfiguration"
	reverseproxy "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/reverseproxy"
	samlsettings "github.com/hmlkao/provider-jfrog-platform/internal/controller/platform/samlsettings"
	providerconfig "github.com/hmlkao/provider-jfrog-platform/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		globalrole.Setup,
		group.Setup,
		httpssosettings.Setup,
		oidcconfiguration.Setup,
		reverseproxy.Setup,
		samlsettings.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
