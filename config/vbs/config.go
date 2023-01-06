// Package vbs contains FlexibleEngine Volume Backup Service (VBS) resources.
package vbs

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vbs_backup_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vbs_backup_v2
	p.AddResourceConfigurator("flexibleengine_vbs_backup_v2", func(r *config.Resource) {

		r.References["volume_id"] = config.Reference{
			Type: tools.GenerateType("evs", "BlockStorageVolume"),
		}

		// snapshot_id
		r.References["snapshot_id"] = config.Reference{
			Type: tools.GenerateType("csbs", "Backup"),
		}
	})

	// flexibleengine_vbs_backup_policy_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vbs_backup_policy_v2
	p.AddResourceConfigurator("flexibleengine_vbs_backup_policy_v2", func(r *config.Resource) {

		// resources
		r.References["resources"] = config.Reference{
			Type:              tools.GenerateType("evs", "BlockStorageVolume"),
			SelectorFieldName: "ResourceSelector",
			RefFieldName:      "ResourceRef",
		}

	})
}
