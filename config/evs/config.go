// Package evs contains custom ResourceConfigurators for FlexibleEngine EVS resources.
package evs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_blockstorage_volume_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/blockstorage_volume_v2
	p.AddResourceConfigurator("flexibleengine_blockstorage_volume_v2", func(r *config.Resource) {

		// snapshot_id
		r.References["snapshot_id"] = config.Reference{
			Type: tools.GenerateType("csbs", "Backup"),
		}

		// metadata.__system__cmkid
		r.References["metadata.__system__cmkid"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}

		// image_id
		r.References["image_id"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}

		// source_replica and source_vol_id require flexibleengine_blockstorage_volume_v2 so we can't reference them

	})

	// flexibleengine_compute_volume_attach_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_volume_attach_v2
	p.AddResourceConfigurator("flexibleengine_compute_volume_attach_v2", func(r *config.Resource) {

		r.References["volume_id"] = config.Reference{
			Type: tools.GenerateType("evs", "BlockStorageVolume"),
		}
	})
}
