package iam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("flexibleengine_identity_group_membership_v3", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			Type: "Group",
		}
		r.References["users"] = config.Reference{
			Type: "User",
		}
	})
	p.AddResourceConfigurator("flexibleengine_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["role_id"] = config.Reference{
			Type: "Role",
		}
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("flexibleengine_identity_provider_conversion", func(r *config.Resource) {
		r.References["provider_id"] = config.Reference{
			Type: "Provider",
		}
	})
}
