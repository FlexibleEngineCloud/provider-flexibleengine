/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	api "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ag/api"
	group "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ag/group"
	apiagd "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/api"
	apipublishment "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/apipublishment"
	application "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/application"
	customauthorizer "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/customauthorizer"
	environment "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/environment"
	groupagd "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/group"
	instance "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/instance"
	response "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/response"
	throttlingpolicy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/throttlingpolicy"
	throttlingpolicyassociate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/throttlingpolicyassociate"
	vpcchannel "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/agd/vpcchannel"
	server "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/bms/server"
	addon "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/addon"
	cluster "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/cluster"
	namespace "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/namespace"
	node "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/node"
	nodepool "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/nodepool"
	pvc "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/cce/pvc"
	alarmrule "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ces/alarmrule"
	clustercss "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/css/cluster"
	snapshot "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/css/snapshot"
	instancedcs "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dcs/instance"
	databaserole "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dds/databaserole"
	databaseuser "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dds/databaseuser"
	instancedds "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dds/instance"
	certificate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/certificate"
	ipgroup "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/ipgroup"
	listener "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/listener"
	loadbalancer "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/loadbalancer"
	member "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/member"
	monitor "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/monitor"
	pool "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dedicatedelb/pool"
	recordset "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dns/recordset"
	zone "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/dns/zone"
	job "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/drs/job"
	floatingip "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/floatingip"
	floatingipassociate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/floatingipassociate"
	instanceecs "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/instance"
	interfaceattach "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/interfaceattach"
	keypair "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/keypair"
	servergroup "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/servergroup"
	volumeattach "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ecs/volumeattach"
	eip "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/eip/eip"
	eipassociate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/eip/eipassociate"
	l7policy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/l7policy"
	l7rule "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/l7rule"
	listenerelb "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/listener"
	loadbalancerelb "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/loadbalancer"
	memberelb "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/member"
	monitorelb "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/monitor"
	poolelb "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/pool"
	whitelist "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/elb/whitelist"
	project "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/eps/project"
	blockstoragevolume "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/evs/blockstoragevolume"
	agency "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/agency"
	groupiam "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/group"
	groupmembership "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/groupmembership"
	projectiam "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/project"
	provider "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/provider"
	providerconversion "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/providerconversion"
	role "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/role"
	roleassignment "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/roleassignment"
	user "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/iam/user"
	image "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/ims/image"
	key "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/kms/key"
	dataset "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/modelarts/dataset"
	datasetversion "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/modelarts/datasetversion"
	dnatrule "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/nat/dnatrule"
	gateway "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/nat/gateway"
	snatrule "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/nat/snatrule"
	obsbucket "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/obsbucket"
	obsbucketobject "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/obsbucketobject"
	obsbucketreplication "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/obsbucketreplication"
	s3bucket "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/s3bucket"
	s3bucketobject "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/s3bucketobject"
	s3bucketpolicy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/oss/s3bucketpolicy"
	providerconfig "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/providerconfig"
	account "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/account"
	database "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/database"
	databaseprivilege "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/databaseprivilege"
	instancerds "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/instance"
	parametergroup "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/parametergroup"
	readreplica "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/rds/readreplica"
	organization "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/swr/organization"
	organizationusers "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/swr/organizationusers"
	repository "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/swr/repository"
	repositorysharing "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/swr/repositorysharing"
	backup "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vbs/backup"
	backuppolicy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vbs/backuppolicy"
	flowlog "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/flowlog"
	network "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/network"
	networkingsubnet "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/networkingsubnet"
	peeringconnection "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/peeringconnection"
	peeringconnectionaccepter "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/peeringconnectionaccepter"
	port "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/port"
	route "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/route"
	router "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/router"
	routerinterface "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/routerinterface"
	routetable "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/routetable"
	secgroup "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/secgroup"
	secgrouprule "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/secgrouprule"
	vip "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/vip"
	vipassociate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/vipassociate"
	vpc "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/vpc"
	vpcsubnet "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpc/vpcsubnet"
	approval "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpcep/approval"
	endpoint "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpcep/endpoint"
	service "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/vpcep/service"
	certificatewaf "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/certificate"
	dedicatedcertificate "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/dedicatedcertificate"
	dedicateddomain "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/dedicateddomain"
	dedicatedinstance "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/dedicatedinstance"
	dedicatedpolicy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/dedicatedpolicy"
	domain "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/domain"
	policy "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/policy"
	rulealarmmasking "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/rulealarmmasking"
	ruleblacklist "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/ruleblacklist"
	ruleccprotection "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/ruleccprotection"
	ruledatamasking "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/ruledatamasking"
	rulepreciseprotection "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/rulepreciseprotection"
	rulewebtamperprotection "github.com/FrangipaneTeam/provider-flexibleengine/internal/controller/waf/rulewebtamperprotection"
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
		alarmrule.Setup,
		clustercss.Setup,
		snapshot.Setup,
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
		floatingip.Setup,
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
