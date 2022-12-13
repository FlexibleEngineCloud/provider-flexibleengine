// Package sfs contains custom ResourceConfigurators for the sfs package.
package sfs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("flexibleengine_rds_instance_v3", func(r *config.Resource) {
		r.References["sfs_id"] = config.Reference{
			Type: "FileSystem",
		}
	})
	p.AddResourceConfigurator("flexibleengine_sfs_turbo", func(r *config.Resource) {
		r.References["vpc_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
		r.References["subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
		r.References["security_group_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecurityGroup"),
		}
		r.References["crypt_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})

}
