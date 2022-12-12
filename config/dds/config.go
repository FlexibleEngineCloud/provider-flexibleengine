package dds

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_dds_instance_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_instance_v3
	p.AddResourceConfigurator("flexibleengine_dds_instance_v3", func(r *config.Resource) {
		// disk_encryption_id
		r.References["disk_encryption_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})

	// flexibleengine_dds_database_user
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_database_user
	p.AddResourceConfigurator("flexibleengine_dds_database_user", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

	// flexibleengine_dds_database_role
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dds_database_role
	p.AddResourceConfigurator("flexibleengine_dds_database_role", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// db_name
		r.References["db_name"] = config.Reference{
			TerraformName: "flexibleengine_dds_database_user",
			Extractor:     common.PathDBNameExtractor,
		}

	})
}
