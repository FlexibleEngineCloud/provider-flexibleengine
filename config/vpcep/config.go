package vpcep

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpcep_service
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_service
	p.AddResourceConfigurator("flexibleengine_vpcep_service", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "EndpointService"

		// vpc_id is the ID of the VPC to which this endpoint service will be attached.
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// port_id is the ID of the port to associate.
		r.References["port_id"] = config.Reference{
			Type: "NetworkPort",
		}

	})

	// flexibleengine_vpcep_endpoint
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_endpoint
	p.AddResourceConfigurator("flexibleengine_vpcep_endpoint", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "Endpoint"

		// vpc_id is the ID of the VPC to which this endpoint will be attached.
		r.References["vpc_id"] = config.Reference{
			Type: "Vpc",
		}

		// network_id is the ID of the network to which this endpoint will be attached.
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}

		// service_id is the ID of the service to which this endpoint will be attached.
		r.References["service_id"] = config.Reference{
			Type: "EndpointService",
		}

	})

	// flexibleengine_vpcep_approval
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_approval
	p.AddResourceConfigurator("flexibleengine_vpcep_approval", func(r *config.Resource) {
		r.ShortGroup = "vpc"
		r.Kind = "EndpointApproval"

		// endpoint_id is the ID of the endpoint to which this approval will be attached.
		r.References["endpoint_id"] = config.Reference{
			Type: "EndPoint",
		}

		// service_id is the ID of the service to which this approval will be attached.
		r.References["service_id"] = config.Reference{
			Type: "EndpointService",
		}

	})

}
