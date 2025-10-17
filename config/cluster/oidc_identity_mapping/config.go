package oidcidentitymapping

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_oidc_identity_mapping", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "OIDCIdentityMapping"
		// Import is supported using the following syntax
		//   terraform import platform_oidc_identity_mapping.my-oidc-identity-mapping my-oidc-identity-mapping:my-oidc-configuration
		//   terraform import platform_oidc_identity_mapping.my-oidc-identity-mapping my-oidc-identity-mapping:my-oidc-configuration:myproj
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			name, _ := tfstate["name"].(string)
			providerName, _ := tfstate["provider_name"].(string)
			projectKey, _ := tfstate["project_key"].(string)
			if name != "" {
				if providerName != "" {
					if projectKey != "" {
						return name + ":" + providerName + ":" + projectKey, nil
					}
					// Return policy id for project if provider_name is set
					return name + ":" + providerName, nil
				}
				// Return policy id
				return "", errors.New("cannot find 'provider_name' in tfstate")
			}
			return "", errors.New("cannot find 'name' in tfstate")
		}
	})
}
