// Package ces provides a set of custom ResourceConfigurators for the CES
package ces

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_ces_alarmrule
	p.AddResourceConfigurator("flexibleengine_ces_alarmrule", func(r *config.Resource) {
		// alarm_actions.notification_list
		r.References["alarm_actions.notification_list"] = config.Reference{
			Type: tools.GenerateType("smn", "Topic"),
		}
	})
}
