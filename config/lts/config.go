package lts

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_lts_topic
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lts_topic
	p.AddResourceConfigurator("flexibleengine_lts_topic", func(r *config.Resource) {

		// vpc_id is the ID of the vpc to which this nat gateway will be attached.
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}

	})

}
