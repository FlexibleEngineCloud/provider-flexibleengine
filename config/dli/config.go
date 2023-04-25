// Package dli contains custom ResourceConfigurators for FlexibleEngine DLI resources.
package dli

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_dli_flinksql_job
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dli_flinksql_job
	p.AddResourceConfigurator("flexibleengine_dli_flinksql_job", func(r *config.Resource) {

		r.References["smn_topic"] = config.Reference{
			Type: tools.GenerateType("smn", "Topic"),
		}

		r.References["obs_bucket"] = config.Reference{
			Type: tools.GenerateType("oss", "OBSBucket"),
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"resume_max_num"},
		}

		// This value generate loop destroy
		config.MarkAsRequired(r.TerraformResource, "resume_max_num")

	})

	p.AddResourceConfigurator("flexibleengine_dli_table", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "Database",
		}
	})

	p.AddResourceConfigurator("flexibleengine_dli_spark_job", func(r *config.Resource) {
		r.References["queue_name"] = config.Reference{
			Type: "Queue",
		}

		r.References["app_name"] = config.Reference{
			Type: "DLIPackage",
		}

		r.References["jars"] = config.Reference{
			Type: "DLIPackage",
		}

		r.References["python_files"] = config.Reference{
			Type: "DLIPackage",
		}

		r.References["files"] = config.Reference{
			Type: "DLIPackage",
		}

	})
}
