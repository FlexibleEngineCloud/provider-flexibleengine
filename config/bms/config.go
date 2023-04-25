// Package bms contains the configuration for the FlexibleEngine BMS service.
package bms

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_compute_bms_server_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_bms_server_v2
	p.AddResourceConfigurator("flexibleengine_compute_bms_server_v2", func(r *config.Resource) {

		// key_pair
		r.References["key_pair"] = config.Reference{
			Type: tools.GenerateType("ecs", "KeyPair"),
		}

		// image_id
		r.References["image_id"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}

		// security_groups
		r.References["security_groups"] = config.Reference{
			TerraformName: "flexibleengine_networking_secgroup_v2",
			Extractor:     tools.GenerateExtractor(false, "name"),
		}

		r.References["network.uuid"] = config.Reference{
			// Require Network ID of VPC Subnet
			TerraformName: "flexibleengine_vpc_subnet_v1",
			Extractor:     tools.GenerateExtractor(true, "id"),
		}

		r.References["network.port"] = config.Reference{
			Type: tools.GenerateType("vpc", "Port"),
		}
	})
}
