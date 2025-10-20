package awsiamrole

import (
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_aws_iam_role", func(r *config.Resource) {
		r.ShortGroup = "platform" // Otherwise 'aws' is used
		r.Kind = "AWSIAMRole"     // Otherwise 'IAMRole' is used
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if username, ok := tfstate["username"].(string); ok && username != "" {
				return username, nil
			}
			return "", errors.New("cannot find 'username' in tfstate")
		}
	})
}
