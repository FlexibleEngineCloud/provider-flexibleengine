/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// compute

	// Imported with the following format : {floating_ip}/{instance_id}/{fixed_ip}
	"flexibleengine_compute_floatingip_associate_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.floating_ip}}/{{.parameters.instance_id}}/{{.parameters.fixed_ip}}"),
	// Imported using the ID
	"flexibleengine_compute_instance_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{port_id}
	"flexibleengine_compute_interface_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.port_id}}"),
	// Imported using name
	"flexibleengine_compute_keypair_v2": config.NameAsIdentifier,
	// Imported using the ID
	"flexibleengine_compute_servergroup_v2": config.IdentifierFromProvider,
	// Imported with the following format : {instance_id}/{volume_id}
	"flexibleengine_compute_volume_attach_v2": TemplatedStringAsIdentifierWithNoName("{{.parameters.instance_id}}/{{.parameters.volume_id}}"),

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
	"flexibleengine_identity_provider_conversion": TemplatedStringAsIdentifierWithNoName("{{.parameters.provider_id}}"),
	// No import documented
	"flexibleengine_identity_role_assignment_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_role_v3": config.IdentifierFromProvider,
	// Imported using the ID
	"flexibleengine_identity_user_v3": config.IdentifierFromProvider,
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
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
