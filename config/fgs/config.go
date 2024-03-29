// Package fgs contains FlexibleEngine FunctionGraph Service (FGS) specific configuration.
package fgs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
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
		// agency
		r.References["agency"] = config.Reference{
			Type:      tools.GenerateType("iam", "Agency"),
			Extractor: tools.GenerateExtractor(false, "name"),
		}

	})

}
