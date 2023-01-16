// Package vpcep contains FlexibleEngine Virtual Private Cloud Endpoint (VPCEP) resources.
package vpcep

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vpcep_service
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_service
	p.AddResourceConfigurator("flexibleengine_vpcep_service", func(r *config.Resource) {
		r.References["port_id"] = config.Reference{
			Type:      tools.GenerateType("ecs", "Instance"),
			Extractor: common.PathNetworkPortIDExtractor,
		}
	})

	// flexibleengine_vpcep_endpoint
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_endpoint
	p.AddResourceConfigurator("flexibleengine_vpcep_endpoint", func(r *config.Resource) {

		// service_id is the ID of the service to which this endpoint will be attached.
		r.References["service_id"] = config.Reference{
			Type: "VPCEPService",
		}

	})

	// flexibleengine_vpcep_approval
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpcep_approval
	p.AddResourceConfigurator("flexibleengine_vpcep_approval", func(r *config.Resource) {

		// endpoints is the ID of the endpoint to which this approval will be attached.
		r.References["endpoints"] = config.Reference{
			Type: "Endpoint",
		}

		// service_id is the ID of the service to which this approval will be attached.
		r.References["service_id"] = config.Reference{
			Type: "VPCEPService",
		}

	})

}
