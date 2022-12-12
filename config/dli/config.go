// Package dli contains custom ResourceConfigurators for FlexibleEngine DLI resources.
package dli

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
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
			Type:      "Database",
			Extractor: common.PathNameExtractor,
		}
	})
}
