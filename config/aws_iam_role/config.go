package awsiamrole

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_aws_iam_role", func(r *config.Resource) {
		r.ShortGroup = "platform"
		r.Kind = "AWSIAMRole"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if username, ok := tfstate["username"].(string); ok && username != "" {
				return username, nil
			}
			return "", errors.New("cannot find 'username' in tfstate")
		}
	})
}
