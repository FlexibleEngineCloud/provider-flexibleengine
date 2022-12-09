/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	floatingipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/floatingipassociate"
	instance "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/instance"
	interfaceattach "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/interfaceattach"
	keypair "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/keypair"
	servergroup "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/servergroup"
	volumeattach "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/volumeattach"
	eip "github.com/gaetanars/provider-flexibleengine/internal/controller/eip/eip"
	eipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/eip/eipassociate"
	agency "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/agency"
	group "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/group"
	groupmembership "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/groupmembership"
	project "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/project"
	provider "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/provider"
	providerconversion "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/providerconversion"
	role "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/role"
	roleassignment "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/roleassignment"
	user "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/user"
	providerconfig "github.com/gaetanars/provider-flexibleengine/internal/controller/providerconfig"
	flowlog "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/flowlog"
	network "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/network"
	networkingsubnet "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/networkingsubnet"
	peeringconnection "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/peeringconnection"
	peeringconnectionaccepter "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/peeringconnectionaccepter"
	port "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/port"
	route "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/route"
	router "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/router"
	routerinterface "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/routerinterface"
	routetable "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/routetable"
	secgroup "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/secgroup"
	secgrouprule "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/secgrouprule"
	vip "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vip"
	vipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vipassociate"
	vpc "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vpc"
	vpcsubnet "github.com/gaetanars/provider-flexibleengine/internal/controller/vpc/vpcsubnet"
	approval "github.com/gaetanars/provider-flexibleengine/internal/controller/vpcep/approval"
	endpoint "github.com/gaetanars/provider-flexibleengine/internal/controller/vpcep/endpoint"
	service "github.com/gaetanars/provider-flexibleengine/internal/controller/vpcep/service"
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
		eip.Setup,
		eipassociate.Setup,
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
		flowlog.Setup,
		network.Setup,
		networkingsubnet.Setup,
		peeringconnection.Setup,
		peeringconnectionaccepter.Setup,
		port.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routetable.Setup,
		secgroup.Setup,
		secgrouprule.Setup,
		vip.Setup,
		vipassociate.Setup,
		vpc.Setup,
		vpcsubnet.Setup,
		approval.Setup,
		endpoint.Setup,
		service.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
