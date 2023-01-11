// Package dms contains custom ResourceConfigurators for the dms package.
package dms

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexengine_dms_kafka_instance
	// https://registry.terraform.io/providers/terraform-provider-flexengine/flexengine/latest/docs/resources/dms_kafka_instance
	p.AddResourceConfigurator("flexengine_dms_kafka_instance", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"available_zones", "access_user", "password", "manager_user", "manager_password"},
		}
	})

}
