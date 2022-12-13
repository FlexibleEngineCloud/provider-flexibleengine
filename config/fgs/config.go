// Package fgs contains FlexibleEngine FunctionGraph Service (FGS) specific configuration.
package fgs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_fgs_trigger
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_trigger
	p.AddResourceConfigurator("flexibleengine_fgs_trigger", func(r *config.Resource) {

		r.References["function_urn"] = config.Reference{
			Type: "Function",
		}

	})

	// flexibleengine_fgs_function
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_function
	p.AddResourceConfigurator("flexibleengine_fgs_function", func(r *config.Resource) {

		r.References["subnet_id"] = config.Reference{
			Type: "VPCSubnet",
		}

		// agency
		r.References["agency"] = config.Reference{
			Type: tools.GenerateType("iam", "Agency"),
		}

	})

}
