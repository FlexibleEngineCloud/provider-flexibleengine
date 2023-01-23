package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{

	// flexibleengine_cce_pvc - Imported using {cluster_id}/{namespace}/{name}
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/cce_pvc
	"flexibleengine_cce_pvc": TemplatedStringAsIdentifierWithNoName("{{ .parameters.cluster_id }}/{{ .parameters.namespace }}/{{ .external_name }}"),

	// flexibleengine_fgs_dependency - Imported using the ID
	// https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/resources/fgs_dependency
	"flexibleengine_fgs_dependency": config.IdentifierFromProvider,
}
