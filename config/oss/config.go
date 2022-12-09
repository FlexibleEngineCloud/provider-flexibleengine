package oss

import "github.com/upbound/upjet/pkg/config"

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

		// TODO This Require KMS

		// kms_key_id is the ID of the KMS key to be used for encryption.
		// r.References["kms_key_id"] = config.Reference{
		// 	Type: "",
		// }

		// TODO This Require s3_bucket
		// Logging target_bucket

	})

	// flexibleengine_obs_bucket_object
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/obs_bucket_object
	p.AddResourceConfigurator("flexibleengine_obs_bucket_object", func(r *config.Resource) {

		// TODO This Require KMS
		// sse_kms_key_id - (Optional) The ID of the kms key. If omitted, the default master key will be used.

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
