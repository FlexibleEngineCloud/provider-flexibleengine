// Package sfs contains custom ResourceConfigurators for the sfs package.
package sfs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_sfs_access_rule_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sfs_access_rule_v2
	p.AddResourceConfigurator("flexibleengine_sfs_access_rule_v2", func(r *config.Resource) {
		r.References["sfs_id"] = config.Reference{
			Type: "FileSystem",
		}
	})
	// flexibleengine_sfs_turbo
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sfs_turbo
	p.AddResourceConfigurator("flexibleengine_sfs_turbo", func(r *config.Resource) {
		// crypt_key_id is a reference to a Key resource
		r.References["crypt_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
		// subnet_id is a reference to a VPC Network ID	resource
		r.References["subnet_id"] = config.Reference{
			Type:      tools.GenerateType("vpc", "VPCSubnet"),
			Extractor: common.PathIDExtractor,
		}
	})

}
