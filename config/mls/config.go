// Package mls provides a set of custom ResourceConfigurators for the MLS
package mls

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_mls_instance_v1
	p.AddResourceConfigurator("flexibleengine_mls_instance_v1", func(r *config.Resource) {
		r.References["network.vpc_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
		r.References["network.subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
		r.References["network.security_group_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
		r.References["mrs_cluster.id"] = config.Reference{
			Type:      tools.GenerateType("mrs", "Cluster"),
			Extractor: tools.GenerateExtractor(true, "id"),
		}
		r.References["mrs_cluster.user_name"] = config.Reference{
			Type:      tools.GenerateType("mrs", "Cluster"),
			Extractor: tools.GenerateExtractor(true, "user_name"),
		}
		r.References["mrs_cluster.user_password"] = config.Reference{
			Type:      tools.GenerateType("mrs", "Cluster"),
			Extractor: tools.GenerateExtractor(true, "user_password"),
		}
	})
}
