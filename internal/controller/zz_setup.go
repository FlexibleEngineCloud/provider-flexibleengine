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
	eip "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/eip"
	eipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/eipassociate"
	endpoint "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/endpoint"
	endpointapproval "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/endpointapproval"
	endpointservice "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/endpointservice"
	flowlogv1 "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/flowlogv1"
	network "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/network"
	networkport "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/networkport"
	networksubnet "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/networksubnet"
	peeringconnection "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/peeringconnection"
	peeringconnectionaccepter "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/peeringconnectionaccepter"
	route "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/route"
	router "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/router"
	routerinterface "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/routerinterface"
	routetable "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/routetable"
	securitygroup "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/securitygrouprule"
	subnet "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/subnet"
	vip "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vip"
	vipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vipassociate"
	vpc "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vpc"
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
		eip.Setup,
		eipassociate.Setup,
		endpoint.Setup,
		endpointapproval.Setup,
		endpointservice.Setup,
		flowlogv1.Setup,
		network.Setup,
		networkport.Setup,
		networksubnet.Setup,
		peeringconnection.Setup,
		peeringconnectionaccepter.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		vip.Setup,
		vipassociate.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
