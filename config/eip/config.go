package eip

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpc_eip
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip
	// * This resource not required configuration

	// flexibleengine_vpc_eip_associate
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip_associate
	p.AddResourceConfigurator("flexibleengine_vpc_eip_associate", func(r *config.Resource) {
		r.Kind = "EIPAssociate"

		// eip_id is the ID of the EIP to associate.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}

		// port_id is the ID of the port to associate.
		r.References["port_id"] = config.Reference{
			Type: "NetworkPort",
		}
	})

}
