// Package smn contains FlexibleEngine Simple Message Notification resources.
package smn

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_smn_subscription_v2
	p.AddResourceConfigurator("flexibleengine_smn_subscription_v2", func(r *config.Resource) {
		r.References["topic_urn"] = config.Reference{
			Type: "Topic",
		}

		// owner Project ID of the topic creator
		r.References["owner"] = config.Reference{
			Type: tools.GenerateType("iam", "Project"),
		}
	})

}
