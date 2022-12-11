/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	api "github.com/gaetanars/provider-flexibleengine/internal/controller/ag/api"
	group "github.com/gaetanars/provider-flexibleengine/internal/controller/ag/group"
	apiagd "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/api"
	apipublishment "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/apipublishment"
	application "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/application"
	customauthorizer "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/customauthorizer"
	environment "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/environment"
	groupagd "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/group"
	instance "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/instance"
	response "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/response"
	throttlingpolicy "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/throttlingpolicy"
	throttlingpolicyassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/throttlingpolicyassociate"
	vpcchannel "github.com/gaetanars/provider-flexibleengine/internal/controller/agd/vpcchannel"
	server "github.com/gaetanars/provider-flexibleengine/internal/controller/bms/server"
	addon "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/addon"
	cluster "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/cluster"
	namespace "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/namespace"
	node "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/node"
	nodepool "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/nodepool"
	pvc "github.com/gaetanars/provider-flexibleengine/internal/controller/cce/pvc"
	instancedcs "github.com/gaetanars/provider-flexibleengine/internal/controller/dcs/instance"
	databaserole "github.com/gaetanars/provider-flexibleengine/internal/controller/dds/databaserole"
	databaseuser "github.com/gaetanars/provider-flexibleengine/internal/controller/dds/databaseuser"
	instancedds "github.com/gaetanars/provider-flexibleengine/internal/controller/dds/instance"
	certificate "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/certificate"
	ipgroup "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/ipgroup"
	listener "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/listener"
	loadbalancer "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/loadbalancer"
	member "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/member"
	monitor "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/monitor"
	pool "github.com/gaetanars/provider-flexibleengine/internal/controller/dedicatedelb/pool"
	recordset "github.com/gaetanars/provider-flexibleengine/internal/controller/dns/recordset"
	zone "github.com/gaetanars/provider-flexibleengine/internal/controller/dns/zone"
	job "github.com/gaetanars/provider-flexibleengine/internal/controller/drs/job"
	floatingipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/floatingipassociate"
	instanceecs "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/instance"
	interfaceattach "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/interfaceattach"
	keypair "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/keypair"
	servergroup "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/servergroup"
	volumeattach "github.com/gaetanars/provider-flexibleengine/internal/controller/ecs/volumeattach"
	eip "github.com/gaetanars/provider-flexibleengine/internal/controller/eip/eip"
	eipassociate "github.com/gaetanars/provider-flexibleengine/internal/controller/eip/eipassociate"
	l7policy "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/l7policy"
	l7rule "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/l7rule"
	listenerelb "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/listener"
	loadbalancerelb "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/loadbalancer"
	memberelb "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/member"
	monitorelb "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/monitor"
	poolelb "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/pool"
	whitelist "github.com/gaetanars/provider-flexibleengine/internal/controller/elb/whitelist"
	project "github.com/gaetanars/provider-flexibleengine/internal/controller/eps/project"
	blockstoragevolume "github.com/gaetanars/provider-flexibleengine/internal/controller/evs/blockstoragevolume"
	agency "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/agency"
	groupiam "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/group"
	groupmembership "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/groupmembership"
	projectiam "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/project"
	provider "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/provider"
	providerconversion "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/providerconversion"
	role "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/role"
	roleassignment "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/roleassignment"
	user "github.com/gaetanars/provider-flexibleengine/internal/controller/iam/user"
	image "github.com/gaetanars/provider-flexibleengine/internal/controller/ims/image"
	key "github.com/gaetanars/provider-flexibleengine/internal/controller/kms/key"
	dataset "github.com/gaetanars/provider-flexibleengine/internal/controller/modelarts/dataset"
	datasetversion "github.com/gaetanars/provider-flexibleengine/internal/controller/modelarts/datasetversion"
	dnatrule "github.com/gaetanars/provider-flexibleengine/internal/controller/nat/dnatrule"
	gateway "github.com/gaetanars/provider-flexibleengine/internal/controller/nat/gateway"
	snatrule "github.com/gaetanars/provider-flexibleengine/internal/controller/nat/snatrule"
	obsbucket "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/obsbucket"
	obsbucketobject "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/obsbucketobject"
	obsbucketreplication "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/obsbucketreplication"
	s3bucket "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/s3bucket"
	s3bucketobject "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/s3bucketobject"
	s3bucketpolicy "github.com/gaetanars/provider-flexibleengine/internal/controller/oss/s3bucketpolicy"
	providerconfig "github.com/gaetanars/provider-flexibleengine/internal/controller/providerconfig"
	account "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/account"
	database "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/database"
	databaseprivilege "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/databaseprivilege"
	instancerds "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/instance"
	parametergroup "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/parametergroup"
	readreplica "github.com/gaetanars/provider-flexibleengine/internal/controller/rds/readreplica"
	organization "github.com/gaetanars/provider-flexibleengine/internal/controller/swr/organization"
	organizationusers "github.com/gaetanars/provider-flexibleengine/internal/controller/swr/organizationusers"
	repository "github.com/gaetanars/provider-flexibleengine/internal/controller/swr/repository"
	repositorysharing "github.com/gaetanars/provider-flexibleengine/internal/controller/swr/repositorysharing"
	backup "github.com/gaetanars/provider-flexibleengine/internal/controller/vbs/backup"
	backuppolicy "github.com/gaetanars/provider-flexibleengine/internal/controller/vbs/backuppolicy"
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
	certificatewaf "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/certificate"
	dedicatedcertificate "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/dedicatedcertificate"
	dedicateddomain "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/dedicateddomain"
	dedicatedinstance "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/dedicatedinstance"
	dedicatedpolicy "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/dedicatedpolicy"
	domain "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/domain"
	policy "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/policy"
	rulealarmmasking "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/rulealarmmasking"
	ruleblacklist "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/ruleblacklist"
	ruleccprotection "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/ruleccprotection"
	ruledatamasking "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/ruledatamasking"
	rulepreciseprotection "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/rulepreciseprotection"
	rulewebtamperprotection "github.com/gaetanars/provider-flexibleengine/internal/controller/waf/rulewebtamperprotection"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		group.Setup,
		apiagd.Setup,
		apipublishment.Setup,
		application.Setup,
		customauthorizer.Setup,
		environment.Setup,
		groupagd.Setup,
		instance.Setup,
		response.Setup,
		throttlingpolicy.Setup,
		throttlingpolicyassociate.Setup,
		vpcchannel.Setup,
		server.Setup,
		addon.Setup,
		cluster.Setup,
		namespace.Setup,
		node.Setup,
		nodepool.Setup,
		pvc.Setup,
		instancedcs.Setup,
		databaserole.Setup,
		databaseuser.Setup,
		instancedds.Setup,
		certificate.Setup,
		ipgroup.Setup,
		listener.Setup,
		loadbalancer.Setup,
		member.Setup,
		monitor.Setup,
		pool.Setup,
		recordset.Setup,
		zone.Setup,
		job.Setup,
		floatingipassociate.Setup,
		instanceecs.Setup,
		interfaceattach.Setup,
		keypair.Setup,
		servergroup.Setup,
		volumeattach.Setup,
		eip.Setup,
		eipassociate.Setup,
		l7policy.Setup,
		l7rule.Setup,
		listenerelb.Setup,
		loadbalancerelb.Setup,
		memberelb.Setup,
		monitorelb.Setup,
		poolelb.Setup,
		whitelist.Setup,
		project.Setup,
		blockstoragevolume.Setup,
		agency.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		projectiam.Setup,
		provider.Setup,
		providerconversion.Setup,
		role.Setup,
		roleassignment.Setup,
		user.Setup,
		image.Setup,
		key.Setup,
		dataset.Setup,
		datasetversion.Setup,
		dnatrule.Setup,
		gateway.Setup,
		snatrule.Setup,
		obsbucket.Setup,
		obsbucketobject.Setup,
		obsbucketreplication.Setup,
		s3bucket.Setup,
		s3bucketobject.Setup,
		s3bucketpolicy.Setup,
		providerconfig.Setup,
		account.Setup,
		database.Setup,
		databaseprivilege.Setup,
		instancerds.Setup,
		parametergroup.Setup,
		readreplica.Setup,
		organization.Setup,
		organizationusers.Setup,
		repository.Setup,
		repositorysharing.Setup,
		backup.Setup,
		backuppolicy.Setup,
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
		certificatewaf.Setup,
		dedicatedcertificate.Setup,
		dedicateddomain.Setup,
		dedicatedinstance.Setup,
		dedicatedpolicy.Setup,
		domain.Setup,
		policy.Setup,
		rulealarmmasking.Setup,
		ruleblacklist.Setup,
		ruleccprotection.Setup,
		ruledatamasking.Setup,
		rulepreciseprotection.Setup,
		rulewebtamperprotection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
