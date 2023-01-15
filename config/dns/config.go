// Package dns contains custom ResourceConfigurators for the dns package.
package dns

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_dns_zone_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_zone_v2
	p.AddResourceConfigurator("flexibleengine_dns_zone_v2", func(r *config.Resource) {
		r.References["router.router_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
	})

	// flexibleengine_dns_recordset_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_recordset_v2
	p.AddResourceConfigurator("flexibleengine_dns_recordset_v2", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})

	// flexibleengine_dns_ptrrecord_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dns_ptrrecord_v2
	p.AddResourceConfigurator("flexibleengine_dns_ptrrecord_v2", func(r *config.Resource) {

		r.References["floatingip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

		// Region is required because the ID of the resource is a combination of the region and the ID.
		config.MarkAsRequired(r.TerraformResource, "region")

	})

}
