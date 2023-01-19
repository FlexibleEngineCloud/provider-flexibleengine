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
		r.UseAsync = true

		// eip_id
		r.References["eip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

		// node_key_pair is sensitive ?
		r.TerraformResource.Schema["node_key_pair"].Sensitive = true
	})

	// flexibleengine_mrs_job_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/mrs_job_v2
	p.AddResourceConfigurator("flexibleengine_mrs_job_v2", func(r *config.Resource) {

		r.UseAsync = true
	})

}
