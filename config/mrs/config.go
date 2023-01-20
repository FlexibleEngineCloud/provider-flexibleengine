// Package mrs contains the configuration for FlexibleEngine MRS resources.
package mrs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_mrs_cluster_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_cluster_v2
	p.AddResourceConfigurator("flexibleengine_mrs_cluster_v2", func(r *config.Resource) {
		// public_ip
		r.References["public_ip"] = config.Reference{
			Type:      tools.GenerateType("eip", "EIP"),
			Extractor: tools.GenerateExtractor(true, "address"),
		}
		// node_key_pair
		r.References["node_key_pair"] = config.Reference{
			Type: tools.GenerateType("ecs", "KeyPair"),
		}
	})

	// flexibleengine_mrs_job_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v2
	p.AddResourceConfigurator("flexibleengine_mrs_job_v2", func(r *config.Resource) {
		// cluster_id
		r.References["cluster_id"] = config.Reference{
			Type: tools.GenerateType("mrs", "Cluster"),
		}
	})

}
