// Package sms contains custom ResourceConfigurators for FlexibleEngine SMS resources.
package sms

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_sms_task
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/sms_task
	p.AddResourceConfigurator("flexibleengine_sms_task", func(r *config.Resource) {
		r.References["vm_template_id"] = config.Reference{
			Type: "ServerTemplate",
		}
	})
}
