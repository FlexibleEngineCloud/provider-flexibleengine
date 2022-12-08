/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	floatingipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/floatingipassociate"
	instance "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/instance"
	interfaceattach "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/interfaceattach"
	keypair "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/keypair"
	servergroup "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/servergroup"
	volumeattach "github.com/gaetanars/provider-flexibleengine/internal/controller/compute/volumeattach"
	agency "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/agency"
	group "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/group"
	groupmembership "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/groupmembership"
	project "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/project"
	provider "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/provider"
	providerconversion "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/providerconversion"
	role "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/role"
	roleassignment "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/roleassignment"
	user "github.com/gaetanars/provider-flexibleengine/internal/controller/identity/user"
	providerconfig "github.com/gaetanars/provider-flexibleengine/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		floatingipassociate.Setup,
		instance.Setup,
		interfaceattach.Setup,
		keypair.Setup,
		servergroup.Setup,
		volumeattach.Setup,
		agency.Setup,
		group.Setup,
		groupmembership.Setup,
		project.Setup,
		provider.Setup,
		providerconversion.Setup,
		role.Setup,
		roleassignment.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
