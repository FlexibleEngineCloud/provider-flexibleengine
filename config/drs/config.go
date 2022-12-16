// Package drs contains custom ResourceConfigurators for the drs package.
package drs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_drs_job
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/drs_job
	p.AddResourceConfigurator("flexibleengine_drs_job", func(r *config.Resource) {

		// db_info.instance_id
		r.References["db_info.instance_id"] = config.Reference{
			Type: tools.GenerateType("rds", "Instance"),
		}

		// db_info.subnet_id
		r.References["db_info.subnet_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}

		r.UseAsync = true
	})

}
