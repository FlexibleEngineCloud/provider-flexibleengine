/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// identity

	// Imported using the ID
	"flexibleengine_identity_agency_v3": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_identity_group_membership_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_group_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_project_v3": config.IdentifierFromProvider,
	// Imported using name
	"flexibleengine_identity_provider": config.NameAsIdentifier,
	// Imported using provider_id
	"flexibleengine_identity_provider_conversion": config.IdentifierFromProvider,
	// No import documented
	"flexibleengine_identity_role_assignment_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_role_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_user_v3": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
