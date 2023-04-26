// Package css contains custom resource configurators for FlexibleEngine CSS resources.
package css

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_css_cluster_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/css_cluster_v1
	p.AddResourceConfigurator("flexibleengine_css_cluster_v1", func(r *config.Resource) {
		r.UseAsync = true

		r.References["node_config.network_info.vpc_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
		r.References["node_config.network_info.subnet_id"] = config.Reference{
			Type:      tools.GenerateType("vpc", "VPCSubnet"),
			Extractor: tools.GenerateExtractor(true, "id"),
		}
		r.References["node_config.network_info.security_group_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
	})

	// flexibleengine_css_snapshot_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/css_snapshot_v1
	p.AddResourceConfigurator("flexibleengine_css_snapshot_v1", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})
}
