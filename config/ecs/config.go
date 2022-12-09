package ecs

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_compute_floatingip_associate_v2", func(r *config.Resource) {
		r.References["floating_ip"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_instance_v2", func(r *config.Resource) {
		r.References["image_id"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}
		r.References["network.uuid"] = config.Reference{
			Type: tools.GenerateType("vpc", "Network"),
		}
		r.References["network.port"] = config.Reference{
			Type: tools.GenerateType("vpc", "Port"),
		}
		r.References["key_pair"] = config.Reference{
			Type: "KeyPair",
		}
		r.References["scheduler_hints.group"] = config.Reference{
			Type: "ServerGroup",
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_interface_attach_v2", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "Network"),
		}
		r.References["port_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "Port"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_keypair_v2", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "public_key")
	})
}
