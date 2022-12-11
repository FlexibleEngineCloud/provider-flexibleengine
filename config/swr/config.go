package swr

import (
	"github.com/gaetanars/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_compute_bms_server_v2
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/compute_bms_server_v2
	p.AddResourceConfigurator("flexibleengine_swr_organization_users", func(r *config.Resource) {

		// organization
		r.References["organization"] = config.Reference{
			Type: "Organization",
		}

		// user_id
		r.References["user_id"] = config.Reference{
			Type: tools.GenerateType("iam", "User"),
		}

	})

	// flexibleengine_swr_repository
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_repository
	p.AddResourceConfigurator("flexibleengine_swr_repository", func(r *config.Resource) {

		// organization
		r.References["organization"] = config.Reference{
			Type: "Organization",
		}

	})

	// flexibleengine_swr_repository_sharing
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/swr_repository_sharing
	p.AddResourceConfigurator("flexibleengine_swr_repository_sharing", func(r *config.Resource) {

		// organization
		r.References["organization"] = config.Reference{
			Type: "Organization",
		}

		// repository
		r.References["repository"] = config.Reference{
			Type: "Repository",
		}

		r.References["repository"] = config.Reference{
			Type: tools.GenerateType("iam", "User"),
		}
	})
}
