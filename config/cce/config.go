// Package cce contains custom resource configurators for FlexibleEngine CCE resources.
package cce

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
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
			Extractor: tools.GenerateExtractor(true, "address"),
		}
	})

	// flexibleengine_cce_namespace
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_namespace
	p.AddResourceConfigurator("flexibleengine_cce_namespace", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
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

	// note(azrod): this resource is affected by a bug in the Upjet
	// The resource fall into tainted state as a result of certain steps in the creation process fail.
	// Please see tainted issue for details https://github.com/upbound/upjet/issues/80

	// p.AddResourceConfigurator("flexibleengine_cce_pvc", func(r *config.Resource) {

	// 	r.References["cluster_id"] = config.Reference{
	// 		Type: "Cluster",
	// 	}

	// 	r.References["namespace"] = config.Reference{
	// 		Type:      "CCENameSpace",
	// 		Extractor: tools.GenerateExtractor(false, "name"),
	// 	}

	// })
}
