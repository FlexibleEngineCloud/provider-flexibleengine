// Package ecs contains FlexibleEngine ECS (Elastic Cloud Server) configuration.
package ecs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_compute_floatingip_associate_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_floatingip_associate_v2
	p.AddResourceConfigurator("flexibleengine_compute_floatingip_associate_v2", func(r *config.Resource) {
		r.References["floating_ip"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
	})

	// flexibleengine_compute_instance_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_instance_v2
	p.AddResourceConfigurator("flexibleengine_compute_instance_v2", func(r *config.Resource) {
		r.References["network.uuid"] = config.Reference{
			Type:      tools.GenerateType("vpc", "VPCSubnet"),
			Extractor: common.PathIDExtractor,
		}

		r.References["network.port"] = config.Reference{
			Type: tools.GenerateType("vpc", "Port"),
		}

		config.MoveToStatus(r.TerraformResource, "network.port")

		r.References["key_pair"] = config.Reference{
			Type: "KeyPair",
		}
		r.References["scheduler_hints.group"] = config.Reference{
			Type: "ServerGroup",
		}
		r.References["image_id"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}
	})

	// flexibleengine_compute_keypair_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_keypair_v2
	p.AddResourceConfigurator("flexibleengine_compute_keypair_v2", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "public_key")
	})

}
