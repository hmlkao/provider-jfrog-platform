/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"platform_aws_iam_role":   config.IdentifierFromProvider,
	"platform_crowd_settings": config.IdentifierFromProvider,
	"platform_global_role":    config.IdentifierFromProvider,
	// Cannot use NameAsIdentifier, 'name' parameter can contain characters which are not allowed in Terraform resource name
	"platform_group":              config.IdentifierFromProvider,
	"platform_http_sso_settings":  config.IdentifierFromProvider,
	"platform_oidc_configuration": config.IdentifierFromProvider,
	"platform_reverse_proxy":      config.IdentifierFromProvider,
	"platform_saml_settings":      config.IdentifierFromProvider,
	"platform_scim_group":         config.IdentifierFromProvider,
	"platform_scim_user":          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
