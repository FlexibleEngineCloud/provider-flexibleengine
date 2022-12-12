// Package nat contains custom ResourceConfigurators for the nat package.
package nat

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_nat_dnat_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/nat_dnat_rule_v2
	p.AddResourceConfigurator("flexibleengine_nat_dnat_rule_v2", func(r *config.Resource) {

		// vpc_id is the ID of the vpc to which this nat gateway will be attached.
		r.References["nat_gateway_id"] = config.Reference{
			Type: "Gateway",
		}
		// subnet_id is the ID of the subnet to which this nat gateway will be attached.
		r.References["floating_ip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

	})

	// flexibleengine_nat_snat_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/nat_snat_rule_v2
	p.AddResourceConfigurator("flexibleengine_nat_snat_rule_v2", func(r *config.Resource) {

		// vpc_id is the ID of the vpc to which this nat gateway will be attached.
		r.References["nat_gateway_id"] = config.Reference{
			Type: "Gateway",
		}
		// subnet_id is the ID of the subnet to which this nat gateway will be attached.
		r.References["floating_ip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

	})

}
