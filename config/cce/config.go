package cce

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_cce_addon_v3", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})
	p.AddResourceConfigurator("flexibleengine_cce_cluster_v3", func(r *config.Resource) {
		r.References["highway_subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "NetworkingSubnet"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_cce_namespace", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		config.MarkAsRequired(r.TerraformResource, "name")
	})
	p.AddResourceConfigurator("flexibleengine_cce_node_pool_v3", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["key_pair"] = config.Reference{
			Type: "KeyPair",
		}
		r.References["data_volumes.kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_cce_node_v3", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["key_pair"] = config.Reference{
			Type: "KeyPair",
		}
		r.References["eip_ids"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
		r.References["data_volumes.kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})
	p.AddResourceConfigurator("flexibleengine_cce_pvc", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["namespace"] = config.Reference{
			Type: "Namespace",
		}
	})
}
