// Package dis contains FlexibleEngine Data Integration Service (DIS) resources.
package dis

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_dis_stream
	p.AddResourceConfigurator("flexibleengine_dis_stream", func(r *config.Resource) {
	})

}
