package evs

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_blockstorage_volume_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/blockstorage_volume_v2
	p.AddResourceConfigurator("flexibleengine_blockstorage_volume_v2", func(r *config.Resource) {

		r.References["image_id"] = config.Reference{
			Type: tools.GenerateType("ims", "Image"),
		}

		// consistency_group_id
		// TODO Add Reference to consistency_group_id
		r.References["consistency_group_id"] = config.Reference{}

		// snapshot_id
		// TODO Add Reference to snapshot_id
		r.References["snapshot_id"] = config.Reference{}

	})
}
