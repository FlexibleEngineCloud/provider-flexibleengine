package tms

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_tms_tags
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/tms_tags
	p.AddResourceConfigurator("flexibleengine_tms_tags", func(r *config.Resource) {
		r.UseAsync = true
	})
}
