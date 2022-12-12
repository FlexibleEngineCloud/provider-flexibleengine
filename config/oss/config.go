package oss

import (
	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

/*

List of all resources in OSS :
	flexibleengine_obs_bucket
	flexibleengine_obs_bucket_object
	flexibleengine_obs_bucket_replication
	flexibleengine_s3_bucket
	flexibleengine_s3_bucket_object
	flexibleengine_s3_bucket_policy

*/

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// flexibleengine_obs_bucket
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket
	p.AddResourceConfigurator("flexibleengine_obs_bucket", func(r *config.Resource) {

		r.References["kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}

		// TODO This Require s3_bucket
		// Logging target_bucket

	})

	// flexibleengine_obs_bucket_object
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_object
	p.AddResourceConfigurator("flexibleengine_obs_bucket_object", func(r *config.Resource) {

		r.References["sse_kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}

	})

	// flexibleengine_obs_bucket_replication
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_replication

	// flexibleengine_s3_bucket
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket

	// flexibleengine_s3_bucket_object
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_object

	// flexibleengine_s3_bucket_policy
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_policy

}
