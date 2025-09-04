package oidcconfiguration

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_oidc_configuration", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "OIDCConfiguration"
		// Import is supported using the following syntax
		//   terraform import platform_oidc_configuration.my-oidc-configuration my-oidc-configuration
		//   terraform import platform_oidc_configuration.my-oidc-configuration my-oidc-configuration:myproj
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			name, _ := tfstate["name"].(string)
			projectKey, _ := tfstate["project_key"].(string)
			if name != "" {
				if projectKey != "" {
					// Return policy id for project if project_key is set
					return name + ":" + projectKey, nil
				}
				// Return policy id
				return name, nil
			}
			return "", errors.New("cannot find 'name' in tfstate")
		}
	})
}
