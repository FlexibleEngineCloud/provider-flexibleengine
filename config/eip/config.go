// Package eip contains custom ResourceConfigurators for the eip package.
package eip

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
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
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"fixed_ip", "port_id", "network_id"},
		}
	})
}
