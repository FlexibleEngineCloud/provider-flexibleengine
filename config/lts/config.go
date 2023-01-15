// Package lts provides a set of custom ResourceConfigurators for the LTS
package lts

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

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

	// flexibleengine_vpc_flow_log_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_flow_log_v1
	p.AddResourceConfigurator("flexibleengine_vpc_flow_log_v1", func(r *config.Resource) {

		r.References["resource_id"] = config.Reference{
			Type:      tools.GenerateType("ecs", "Instance"),
			Extractor: common.PathNetworkPortIDExtractor,
		}

		r.References["log_group_id"] = config.Reference{
			Type: tools.GenerateType("lts", "Group"),
		}

		r.References["log_topic_id"] = config.Reference{
			Type: tools.GenerateType("lts", "Topic"),
		}

	})

}
