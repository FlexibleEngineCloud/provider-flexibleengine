// Package dms contains custom ResourceConfigurators for the dms package.
package dms

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_dms_rocketmq_instance
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dms_rocketmq_instance
	p.AddResourceConfigurator("flexibleengine_dms_rocketmq_instance", func(r *config.Resource) {
		// publicip_id
		r.References["publicip_id"] = config.Reference{
			Type: tools.GenerateType("eip", "EIP"),
		}
	})

	// flexibleengine_dms_rocketmq_user
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/dms_rocketmq_user
	p.AddResourceConfigurator("flexibleengine_dms_rocketmq_user", func(r *config.Resource) {
		// topic_perms.name
		r.References["topic_perms.name"] = config.Reference{
			Type: "RocketMQTopic",
		}
		// group_perms.name
		r.References["group_perms.name"] = config.Reference{
			Type: "RocketMQConsumerGroup",
		}
	})

}
