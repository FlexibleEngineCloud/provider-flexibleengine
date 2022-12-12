// Package ag contains FlexibleEngine API Gateway (AG) resources.
package ag

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_api_gateway_api
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/api_gateway_api
	p.AddResourceConfigurator("flexibleengine_api_gateway_api", func(r *config.Resource) {

		// group_id
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}

		r.UseAsync = true

	})

}
