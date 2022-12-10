package iam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_identity_group_membership_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_group_membership_v3
	p.AddResourceConfigurator("flexibleengine_identity_group_membership_v3", func(r *config.Resource) {
		r.References["group"] = config.Reference{
			Type: "Group",
		}
		r.References["users"] = config.Reference{
			Type: "User",
		}
	})

	// flexibleengine_identity_role_assignment_v3
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_role_assignment_v3
	p.AddResourceConfigurator("flexibleengine_identity_role_assignment_v3", func(r *config.Resource) {
		r.References["role_id"] = config.Reference{
			Type: "Role",
		}
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})

	// flexibleengine_identity_provider_conversion
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/identity_provider_conversion
	p.AddResourceConfigurator("flexibleengine_identity_provider_conversion", func(r *config.Resource) {
		r.References["provider_id"] = config.Reference{
			Type: "Provider",
		}
	})
}
