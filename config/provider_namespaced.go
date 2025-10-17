/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	awsiamrole "github.com/hmlkao/provider-jfrog-platform/config/namespaced/aws_iam_role"
	crowdsettings "github.com/hmlkao/provider-jfrog-platform/config/namespaced/crowd_settings"
	globalrole "github.com/hmlkao/provider-jfrog-platform/config/namespaced/global_role"
	"github.com/hmlkao/provider-jfrog-platform/config/namespaced/group"
	httpssosettings "github.com/hmlkao/provider-jfrog-platform/config/namespaced/http_sso_settings"
	oidcconfiguration "github.com/hmlkao/provider-jfrog-platform/config/namespaced/oidc_configuration"
	"github.com/hmlkao/provider-jfrog-platform/config/namespaced/permission"
	reverseproxy "github.com/hmlkao/provider-jfrog-platform/config/namespaced/reverse_proxy"
	samlsettings "github.com/hmlkao/provider-jfrog-platform/config/namespaced/saml_settings"
	scimgroup "github.com/hmlkao/provider-jfrog-platform/config/namespaced/scim_group"
	scimuser "github.com/hmlkao/provider-jfrog-platform/config/namespaced/scim_user"
	workersservice "github.com/hmlkao/provider-jfrog-platform/config/namespaced/workers_service"
)

// GetProvider returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("jfrog.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		awsiamrole.Configure,
		crowdsettings.Configure,
		globalrole.Configure,
		group.Configure,
		httpssosettings.Configure,
		oidcconfiguration.Configure,
		permission.Configure,
		reverseproxy.Configure,
		samlsettings.Configure,
		scimgroup.Configure,
		scimuser.Configure,
		workersservice.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
