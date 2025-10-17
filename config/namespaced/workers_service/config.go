package workersservice

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_workers_service", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "WorkersService"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if key, ok := tfstate["key"].(string); ok && key != "" {
				return key, nil
			}
			return "", errors.New("cannot find 'key' in tfstate")
		}
	})
}
