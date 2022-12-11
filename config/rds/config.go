package rds

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_rds_instance_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_instance_v3
	p.AddResourceConfigurator("flexibleengine_rds_instance_v3", func(r *config.Resource) {

		// disk_encryption_id
		r.References["volume.disk_encryption_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})

	// flexibleengine_rds_account
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_account
	p.AddResourceConfigurator("flexibleengine_rds_account", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// TODO password
	})

	// flexibleengine_rds_read_replica_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_read_replica_v3
	p.AddResourceConfigurator("flexibleengine_rds_read_replica_v3", func(r *config.Resource) {

		// instance_id
		r.References["replica_of_id"] = config.Reference{
			Type: "Instance",
		}

		// disk_encryption_id
		r.References["volume.disk_encryption_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
	})

	// flexibleengine_rds_database
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_database
	p.AddResourceConfigurator("flexibleengine_rds_database", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	// flexibleengine_rds_database_privilege
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rds_database_privilege
	p.AddResourceConfigurator("flexibleengine_rds_database_privilege", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

}
