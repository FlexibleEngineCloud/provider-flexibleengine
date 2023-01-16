// Package ecs contains FlexibleEngine ECS (Elastic Cloud Server) configuration.
package ecs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_compute_floatingip_associate_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_floatingip_associate_v2
	p.AddResourceConfigurator("flexibleengine_compute_floatingip_associate_v2", func(r *config.Resource) {
		r.References["floating_ip"] = config.Reference{
			Type:      tools.GenerateType("eip", "EIP"),
			Extractor: common.PathAddressExtractor,
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

	// flexibleengine_compute_interface_attach_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_interface_attach_v2
	p.AddResourceConfigurator("flexibleengine_compute_interface_attach_v2", func(r *config.Resource) {

		// note(azrod): This is not working because if resource has a reference
		// Terraform schema optional is forced to true
		// https://github.com/upbound/upjet/blob/84e87589ebcfadd2b703bae6b916e1c0d42c5b3d/pkg/types/field.go#L184-L195
		//
		// port_id is required because it is used in the import function
		// config.MarkAsRequired(r.TerraformResource, "port_id")
		r.MetaResource.ArgumentDocs["port_id"] = "- (Required) The ID of the Port to attach to an Instance."

		// network_id and fixed_ip are not required because they are not
		// compatible if port_id is defined but they are used in the status
		config.MoveToStatus(r.TerraformResource, "network_id")
		r.MetaResource.ArgumentDocs["network_id"] = "The ID of the Network associated to the port."
		config.MoveToStatus(r.TerraformResource, "fixed_ip")
		r.MetaResource.ArgumentDocs["fixed_ip"] = "The IP address associated to the port."

	})

}
