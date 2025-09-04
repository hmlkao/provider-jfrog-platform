/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal jfrog-platform credentials as JSON"
)

var reqFields = []string{}
var optFields = []string{"url", "access_token", "tfc_credential_tag_name", "oidc_provider_name"}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		providerConfig, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		tfmProviderSetup := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		data, err := resource.CommonCredentialExtractor(ctx, providerConfig.Spec.Credentials.Source, client, providerConfig.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return tfmProviderSetup, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return tfmProviderSetup, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		tfmProviderSetup.Configuration = map[string]any{}
		// Required fields
		for _, req := range reqFields {
			tfmProviderSetup.Configuration[req] = creds[req]
		}
		// Optional fields
		for _, opt := range optFields {
			if v, ok := creds[opt]; ok {
				tfmProviderSetup.Configuration[opt] = v
			}
		}

		return tfmProviderSetup, nil
	}
}
