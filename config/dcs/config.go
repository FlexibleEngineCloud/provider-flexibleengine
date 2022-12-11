package dcs

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_dcs_instance_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dcs_instance_v1
	p.AddResourceConfigurator("flexibleengine_dcs_instance_v1", func(r *config.Resource) {

		// password
		// TODO Secure password

		r.UseAsync = true

	})

}
