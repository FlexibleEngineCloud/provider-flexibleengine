// Package agd contains the configuration for the FlexibleEngine provider.
package agd

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_apig_api
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_api
	p.AddResourceConfigurator("flexibleengine_apig_api", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// group_id
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}

		// authorizer_id
		// TODO authorizer_id

		// web_policy.authorizer_id
		// TODO web_policy.authorizer_id

	})

	// flexibleengine_apig_throttling_policy_associate
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_throttling_policy_associate
	p.AddResourceConfigurator("flexibleengine_apig_throttling_policy_associate", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// policy_id
		r.References["policy_id"] = config.Reference{
			Type: "ThrottlingPolicy",
		}

	})

	// flexibleengine_apig_environment
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_environment
	p.AddResourceConfigurator("flexibleengine_apig_environment", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

	// flexibleengine_apig_vpc_channel
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_vpc_channel
	p.AddResourceConfigurator("flexibleengine_apig_vpc_channel", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

	// flexibleengine_apig_application
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_application
	p.AddResourceConfigurator("flexibleengine_apig_application", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

	// flexibleengine_apig_response
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_response
	p.AddResourceConfigurator("flexibleengine_apig_response", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// api_id
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}

	})

	// flexibleengine_apig_throttling_policy
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_throttling_policy
	p.AddResourceConfigurator("flexibleengine_apig_throttling_policy", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// app_throttles.throttling_object_id
		r.References["app_throttles.throttling_object_id"] = config.Reference{
			Type: "Application",
		}

	})

	// flexibleengine_apig_instance
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_instance
	p.AddResourceConfigurator("flexibleengine_apig_instance", func(r *config.Resource) {

		// eip_id
		r.References["eip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}

	})

	// flexibleengine_apig_group
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_group
	p.AddResourceConfigurator("flexibleengine_apig_group", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// environnement.environnement_id
		r.References["environnement.environnement_id"] = config.Reference{
			Type: "Environment",
		}

	})

	// flexibleengine_apig_custom_authorizer
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_custom_authorizer
	p.AddResourceConfigurator("flexibleengine_apig_custom_authorizer", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

	})

	// flexibleengine_apig_api_publishment
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/apig_api_publishment
	p.AddResourceConfigurator("flexibleengine_apig_api_publishment", func(r *config.Resource) {

		// instance_id
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}

		// api_id
		r.References["api_id"] = config.Reference{
			Type: "API",
		}

		// env_id
		r.References["env_id"] = config.Reference{
			Type: "Environment",
		}

	})

}
