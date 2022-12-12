package bms

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_compute_bms_server_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_bms_server_v2
	p.AddResourceConfigurator("flexibleengine_blockstorage_volume_v2", func(r *config.Resource) {

		// security_groups
		// TODO Here security_groups require name instead of id
		// r.References["security_groups"] = config.Reference{}

		// network.uuid
		r.References["network.uuid"] = config.Reference{
			Type: tools.GenerateType("vpc", "Network"),
		}

		// network.port
		r.References["network.port"] = config.Reference{
			Type: tools.GenerateType("vpc", "NetworkPort"),
		}
	})
}
