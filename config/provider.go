/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/hmlkao/provider-jfrog-platform/config/group"
	reverseproxy "github.com/hmlkao/provider-jfrog-platform/config/reverse_proxy"
	samlsettings "github.com/hmlkao/provider-jfrog-platform/config/saml_settings"
)

const (
	resourcePrefix = "jfrog-platform"
	modulePath     = "github.com/hmlkao/provider-jfrog-platform"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

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
		group.Configure,
		reverseproxy.Configure,
		samlsettings.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
