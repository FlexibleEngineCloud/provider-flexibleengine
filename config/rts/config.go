// Package rts provides a set of custom ResourceConfigurators for the RTS
package rts

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// rts_stack_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/rts_stack_v1
	p.AddResourceConfigurator("flexibleengine_rts_stack_v1", func(r *config.Resource) {
		r.UseAsync = true
	})

}
