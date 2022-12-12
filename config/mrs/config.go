package mrs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_mrs_cluster_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_cluster_v1
	p.AddResourceConfigurator("flexibleengine_mrs_cluster_v1", func(r *config.Resource) {
		// TODO: Add references to password_of_mrs_manager

	})

	// flexibleengine_mrs_cluster_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_cluster_v2
	p.AddResourceConfigurator("flexibleengine_mrs_cluster_v2", func(r *config.Resource) {
		r.UseAsync = true

		// eip_id
		r.References["eip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["manager_admin_pwd"].(string); ok {
				conn["manager_admin_pwd"] = []byte(a)
			}

			if a, ok := attr["node_key_pair"].(string); ok {
				conn["node_key_pair"] = []byte(a)
			}
			return conn, nil
		}
	})

	// flexibleengine_mrs_hybrid_cluster_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_hybrid_cluster_v1
	p.AddResourceConfigurator("flexibleengine_mrs_hybrid_cluster_v1", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["cluster_admin_secret"].(string); ok {
				conn["mrs_hcluster_admin_secret"] = []byte(a)
			}

			if a, ok := attr["master_node_key_pair"].(string); ok {
				conn["mrs_hcluster_master_node_key_pair"] = []byte(a)
			}
			return conn, nil
		}
	})

	// flexibleengine_mrs_job_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v1
	p.AddResourceConfigurator("flexibleengine_mrs_job_v1", func(r *config.Resource) {

	})

	// flexibleengine_mrs_job_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v2
	p.AddResourceConfigurator("flexibleengine_mrs_job_v2", func(r *config.Resource) {
	})

}
