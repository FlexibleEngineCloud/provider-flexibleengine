// Package cse provides a set of custom ResourceConfigurators for the CSE
package cse

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_cse_microservice_engine
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice_engine
	p.AddResourceConfigurator("flexibleengine_cse_microservice_engine", func(r *config.Resource) {

		r.UseAsync = true

		r.References["eip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

	})

	// flexibleengine_cse_microservice
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice
	p.AddResourceConfigurator("flexibleengine_cse_microservice", func(r *config.Resource) {

		r.UseAsync = true

		// connect_address
		r.References["connect_address"] = config.Reference{
			Type: "MicroserviceEngine",
		}

	})

	// flexibleengine_cse_microservice_instance
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cse_microservice_instance
	p.AddResourceConfigurator("flexibleengine_cse_microservice_instance", func(r *config.Resource) {

		r.UseAsync = true

		// connect_address
		r.References["connect_address"] = config.Reference{
			Type: "MicroserviceEngine",
		}

		// microservice_id
		r.References["microservice_id"] = config.Reference{
			Type: "Microservice",
		}

	})

}
