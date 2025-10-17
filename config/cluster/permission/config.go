package permission

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_permission", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "Permission"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if name, ok := tfstate["name"].(string); ok && name != "" {
				return name, nil
			}
			return "", errors.New("cannot find 'name' in tfstate")
		}
	})
}
