// Package antiddos provides a flexibleengine_antiddos_v1 ResourceConfigurator.
package antiddos

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_antiddos_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/antiddos_v1
	p.AddResourceConfigurator("flexibleengine_antiddos_v1", func(r *config.Resource) {

		r.References["floating_ip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

	})

}
