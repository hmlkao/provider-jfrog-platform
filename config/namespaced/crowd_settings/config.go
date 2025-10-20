package crowdsettings

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_crowd_settings", func(r *config.Resource) {
		r.ShortGroup = "platform" // Otherwise 'crowd' is used
		r.Kind = "CrowdSettings"  // Otherwise 'Settings' is used
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if applicationName, ok := tfstate["application_name"].(string); ok && applicationName != "" {
				return applicationName, nil
			}
			return "", errors.New("cannot find 'application_name' in tfstate")
		}
	})
}
