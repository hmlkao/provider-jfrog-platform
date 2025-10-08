/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	awsiamrole "github.com/hmlkao/provider-jfrog-platform/config/cluster/aws_iam_role"
	crowdsettings "github.com/hmlkao/provider-jfrog-platform/config/cluster/crowd_settings"
	globalrole "github.com/hmlkao/provider-jfrog-platform/config/cluster/global_role"
	"github.com/hmlkao/provider-jfrog-platform/config/cluster/group"
	httpssosettings "github.com/hmlkao/provider-jfrog-platform/config/cluster/http_sso_settings"
	oidcconfiguration "github.com/hmlkao/provider-jfrog-platform/config/cluster/oidc_configuration"
	reverseproxy "github.com/hmlkao/provider-jfrog-platform/config/cluster/reverse_proxy"
	samlsettings "github.com/hmlkao/provider-jfrog-platform/config/cluster/saml_settings"
	scimgroup "github.com/hmlkao/provider-jfrog-platform/config/cluster/scim_group"
)

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("jfrog.crossplane.io"),
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
		reverseproxy.Configure,
		samlsettings.Configure,
		scimgroup.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
