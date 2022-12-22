// Package lts provides a set of custom ResourceConfigurators for the LTS
package lts

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_lts_topic
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/lts_topic
	p.AddResourceConfigurator("flexibleengine_lts_topic", func(r *config.Resource) {

		// group_id is a reference to a Group
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}

	})

}
