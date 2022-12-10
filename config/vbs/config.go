package vbs

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_vbs_backup_v2
	p.AddResourceConfigurator("flexibleengine_vbs_backup_v2", func(r *config.Resource) {

		r.References["volume_id"] = config.Reference{
			Type: tools.GenerateType("evs", "BlockStorageVolume"),
		}

		// snapshot_id
		// TODO Add Reference to snapshot_id
		r.References["snapshot_id"] = config.Reference{}

	})
}
