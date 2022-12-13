package cts

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_cts_tracker_v1
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cts_tracker_v1
	p.AddResourceConfigurator("flexibleengine_cts_tracker_v1", func(r *config.Resource) {
		// bucket_name
		r.References["bucket_name"] = config.Reference{
			Type:      tools.GenerateType("oss", "OBSBucket"),
			Extractor: common.PathBucketNameExtractor,
		}
	})
}
