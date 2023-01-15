// Package waf contains the configuration for the FlexibleEngine WAF service.
package waf

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_waf_dedicated_domain
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_domain
	p.AddResourceConfigurator("flexibleengine_waf_dedicated_domain", func(r *config.Resource) {

		r.References["certificate_id"] = config.Reference{
			Type: "DedicatedCertificate",
		}

		r.References["policy_id"] = config.Reference{
			Type: "DedicatedPolicy",
		}

		r.References["server.vcp_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
	})

	// flexibleengine_waf_dedicated_instance
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_dedicated_instance
	p.AddResourceConfigurator("flexibleengine_waf_dedicated_instance", func(r *config.Resource) {

		r.References["security_group"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}

		r.References["group_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "ServerGroup"),
		}
	})

	// flexibleengine_waf_domain
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/waf_domain
	p.AddResourceConfigurator("flexibleengine_waf_domain", func(r *config.Resource) {

		r.References["certificate_id"] = config.Reference{
			Type: "Certificate",
		}

	})

}
