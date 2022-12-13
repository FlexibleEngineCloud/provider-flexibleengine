// Package csbs provides a Configurator for the Crossplane Service Broker Service (CSBS).
package csbs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_csbs_backup_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/csbs_backup_v1
	p.AddResourceConfigurator("flexibleengine_csbs_backup_v1", func(r *config.Resource) {
		r.References["resource_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "Instance"),
		}
		r.UseAsync = true
	})
}
