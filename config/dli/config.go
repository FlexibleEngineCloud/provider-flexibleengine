// Package dli contains custom ResourceConfigurators for FlexibleEngine DLI resources.
package dli

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("flexibleengine_dli_flinksql_job", func(r *config.Resource) {
		r.References["smn_topic"] = config.Reference{
			Type: tools.GenerateType("smn", "Topic"),
		}
	})

	p.AddResourceConfigurator("flexibleengine_dli_table", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "Database",
		}
	})

	// flexibleengine_dli_spark_job
	p.AddResourceConfigurator("flexibleengine_dli_spark_job", func(r *config.Resource) {
		r.References["queue_name"] = config.Reference{
			Type: "Queue",
		}

		// app_name
		r.References["app_name"] = config.Reference{
			Type: "DLIPackage",
		}

		// jars
		r.References["jars"] = config.Reference{
			Type: "DLIPackage",
		}

		// python_files
		r.References["python_files"] = config.Reference{
			Type: "DLIPackage",
		}

		// files
		r.References["files"] = config.Reference{
			Type: "DLIPackage",
		}

	})
}
