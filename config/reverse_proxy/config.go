package reverseproxy

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_reverse_proxy", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "ReverseProxy"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if serverProvider, ok := tfstate["server_provider"].(string); ok && serverProvider != "" {
				return serverProvider, nil
			}
			return "", errors.New("cannot find 'server_provider' in tfstate")
		}
	})
}
