package group

import (
	"github.com/crossplane/upjet/pkg/config"
)

// const shortGroup string = "" // Otherwise, 'local' is used

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_group", func(r *config.Resource) {
		// r.ShortGroup = shortGroup
		// Specify Kubernetes kind, otherwise, 'AnsibleRepository' is used
		// r.Kind = "LocalAnsibleRepository"
		// Set custom func to get external name because there is no 'id' stored in Terraform state
	})
}
