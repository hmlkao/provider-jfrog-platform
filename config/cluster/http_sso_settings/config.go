package httpssosettings

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_http_sso_settings", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "HTTPSSOSettings"
		// Cannot use config.NameAsIdentifier variable because 'name' parameter can use characters which are invalid for Trerraform resource name
		//   Terraform resource name must start with a letter or underscore and may contain only letters, digits, underscores, and dashes.
		// Variable config.NameAsIdentifier is using IDAsExternalName func which tries to get the 'id' from the tfstate,
		// but there is no 'id' after the Terraform state refresh, so we specify custum function to get 'name'
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if name, ok := tfstate["name"].(string); ok && name != "" {
				return name, nil
			}
			return "", errors.New("cannot find 'name' in tfstate")
		}
	})
}
