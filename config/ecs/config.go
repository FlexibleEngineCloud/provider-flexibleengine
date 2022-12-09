package ecs

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_compute_floatingip_associate_v2", func(r *config.Resource) {
		r.References["floating_ip"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/eip/v1beta1.EIP",
		}
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_instance_v2", func(r *config.Resource) {
		// r.References["image_id"] = config.Reference{
		// 	Type: "github.com/gaetanars/provider-flexibleengine/apis/images/v1beta1.Image",
		// }
		r.References["network.uuid"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network",
		}
		r.References["network.port"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Port",
		}
		r.References["key_pair"] = config.Reference{
			Type: "KeyPair",
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_interface_attach_v2", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["network_id"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network",
		}
		r.References["port_id"] = config.Reference{
			Type: "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Port",
		}
	})
	p.AddResourceConfigurator("flexibleengine_compute_keypair_v2", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "public_key")
	})
}
