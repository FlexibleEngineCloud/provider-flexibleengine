// Package eip contains custom ResourceConfigurators for the eip package.
package eip

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_vpc_eip_associate
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/vpc_eip_associate
	p.AddResourceConfigurator("flexibleengine_vpc_eip_associate", func(r *config.Resource) {
		r.References["public_ip"] = config.Reference{
			Type:      "EIP",
			Extractor: common.PathAddressExtractor,
		}
		r.References["netwok_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPCSubnet"),
		}
	})
}
