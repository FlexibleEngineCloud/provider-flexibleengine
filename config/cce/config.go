// Package cce contains custom resource configurators for FlexibleEngine CCE resources.
package cce

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_cce_addon_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_addon_v3
	p.AddResourceConfigurator("flexibleengine_cce_addon_v3", func(r *config.Resource) {

		r.UseAsync = true

		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})

	// flexibleengine_cce_cluster_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_cluster_v3
	p.AddResourceConfigurator("flexibleengine_cce_cluster_v3", func(r *config.Resource) {
		r.References["highway_subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
		r.References["eip"] = config.Reference{
			Type:      tools.GenerateType("eip", "EIP"),
			Extractor: common.PathAddressExtractor,
		}
	})

	// flexibleengine_cce_namespace
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_namespace
	p.AddResourceConfigurator("flexibleengine_cce_namespace", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		//config.MarkAsRequired(r.TerraformResource, "name")
	})

	// flexibleengine_cce_node_pool_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_node_pool_v3
	p.AddResourceConfigurator("flexibleengine_cce_node_pool_v3", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["key_pair"] = config.Reference{
			Type: tools.GenerateType("ecs", "KeyPair"),
		}
		r.References["data_volumes.kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
		r.References["subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
	})

	// flexibleengine_cce_node_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_node_v3
	p.AddResourceConfigurator("flexibleengine_cce_node_v3", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["key_pair"] = config.Reference{
			Type: tools.GenerateType("ecs", "KeyPair"),
		}
		r.References["eip_ids"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
		r.References["data_volumes.kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})

	// flexibleengine_cce_pvc
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_pvc
	p.AddResourceConfigurator("flexibleengine_cce_pvc", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		// Rename to avoid conflict with the Namespace Kind of Kubernetes
		r.References["namespace"] = config.Reference{
			Type: "CCENameSpace",
		}
	})
}
