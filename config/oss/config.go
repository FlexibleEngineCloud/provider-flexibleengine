// Package oss provides a set of custom ResourceConfigurators for the OSS
package oss

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
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
	})

	// flexibleengine_obs_bucket_object
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_object
	p.AddResourceConfigurator("flexibleengine_obs_bucket_object", func(r *config.Resource) {

		r.References["sse_kms_key_id"] = config.Reference{
			Type: tools.GenerateType("kms", "Key"),
		}
		r.References["bucket"] = config.Reference{
			Type: "OBSBucket",
		}

	})

	// flexibleengine_obs_bucket_replication
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_replication
	p.AddResourceConfigurator("flexibleengine_obs_bucket_replication", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			Type: "OBSBucket",
		}
		r.References["destination_bucket"] = config.Reference{
			Type: "OBSBucket",
		}
		r.References["agency"] = config.Reference{
			Type:      tools.GenerateType("iam", "Agency"),
			Extractor: tools.GenerateExtractor(false, "name"),
		}
	})

	// flexibleengine_s3_bucket
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket

	// flexibleengine_s3_bucket_object
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_object
	p.AddResourceConfigurator("flexibleengine_s3_bucket_object", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			Type: "S3Bucket",
		}
	})

	// flexibleengine_s3_bucket_policy
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/s3_bucket_policy
	p.AddResourceConfigurator("flexibleengine_s3_bucket_policy", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			Type: "S3Bucket",
		}
	})

}
