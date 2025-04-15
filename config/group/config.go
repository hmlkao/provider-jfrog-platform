package group

import (
	"github.com/crossplane/upjet/pkg/config"
)

// const shortGroup string = "" // Otherwise, 'groups' is used

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("platform_group", func(r *config.Resource) {
		// r.ShortGroup = shortGroup
	})
}
