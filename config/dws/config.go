// Package dws contains custom resource configurators for DWS resources.
package dws

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_dws_cluster_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dws_cluster_v1
	p.AddResourceConfigurator("flexibleengine_dws_cluster_v1", func(r *config.Resource) {
		// user_name is sensitive
		r.TerraformResource.Schema["user_name"].Sensitive = true
		// security_group_id
		r.References["security_group_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
		// public_ip.eip_id
		r.References["public_ip.eip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
	})

}
