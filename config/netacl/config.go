// Package netacl contains the configuration for the netacl package.
package netacl

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_network_acl
	p.AddResourceConfigurator("flexibleengine_network_acl", func(r *config.Resource) {
		r.References["inbound_rules"] = config.Reference{
			Type:              tools.GenerateType("netacl", "ACLRule"),
			SelectorFieldName: "InboundRuleSelector",
			RefFieldName:      "InboundRuleRefs",
		}
		r.References["outbound_rules"] = config.Reference{
			Type:              tools.GenerateType("netacl", "ACLRule"),
			SelectorFieldName: "OutboundRuleSelector",
			RefFieldName:      "OutboundRuleRefs",
		}
	})
}
