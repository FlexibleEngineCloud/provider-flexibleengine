// Package waf contains the configuration for the FlexibleEngine WAF service.
package waf

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("flexibleengine_waf_dedicated_domain", func(r *config.Resource) {

		r.References["certificate_id"] = config.Reference{
			Type: "DedicatedCertificate",
		}

		r.References["policy_id"] = config.Reference{
			Type: "DedicatedPolicy",
		}

		r.References["server.vcp_id"] = config.Reference{
			Type: tools.GenerateType("vpc", "VPC"),
		}
	})

	// flexibleengine_waf_dedicated_instance
	p.AddResourceConfigurator("flexibleengine_waf_dedicated_instance", func(r *config.Resource) {

		r.References["security_group"] = config.Reference{
			Type: tools.GenerateType("vpc", "SecGroup"),
		}

		r.References["group_id"] = config.Reference{
			Type: tools.GenerateType("ecs", "ServerGroup"),
		}
	})

	// flexibleengine_waf_domain
	p.AddResourceConfigurator("flexibleengine_waf_domain", func(r *config.Resource) {

		r.References["certificate_id"] = config.Reference{
			Type: "Certificate",
		}

		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}

	})

	// flexibleengine_waf_rule_alarm_masking
	p.AddResourceConfigurator("flexibleengine_waf_rule_alarm_masking", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}

	})

	// flexibleengine_waf_rule_blacklist
	p.AddResourceConfigurator("flexibleengine_waf_rule_blacklist", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}
	})

	// flexibleengine_waf_rule_cc_protection
	p.AddResourceConfigurator("flexibleengine_waf_rule_cc_protection", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}
	})

	// flexibleengine_waf_rule_data_masking
	p.AddResourceConfigurator("flexibleengine_waf_rule_data_masking", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}
	})

	// flexibleengine_waf_rule_precise_protection
	p.AddResourceConfigurator("flexibleengine_waf_rule_precise_protection", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}
	})

	// flexibleengine_waf_rule_web_tamper_protection
	p.AddResourceConfigurator("flexibleengine_waf_rule_web_tamper_protection", func(r *config.Resource) {

		// ! Require MultiType Reference
		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}
	})

}
