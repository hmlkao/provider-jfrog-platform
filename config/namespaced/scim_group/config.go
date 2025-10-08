package scimgroup

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_scim_group", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "SCIMGroup"
		// r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		// 	if displayName, ok := tfstate["display_name"].(string); ok && displayName != "" {
		// 		return displayName, nil
		// 	}
		// 	return "", errors.New("cannot find 'display_name' in tfstate")
		// }
	})
}
