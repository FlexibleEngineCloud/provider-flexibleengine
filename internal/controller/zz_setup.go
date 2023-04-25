/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	api "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ag/api"
	group "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ag/group"
	apiagd "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/api"
	apipublishment "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/apipublishment"
	application "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/application"
	customauthorizer "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/customauthorizer"
	environment "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/environment"
	groupagd "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/group"
	instance "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/instance"
	response "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/response"
	throttlingpolicy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/throttlingpolicy"
	throttlingpolicyassociate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/throttlingpolicyassociate"
	vpcchannel "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/agd/vpcchannel"
	antiddos "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/antiddos/antiddos"
	configuration "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/as/configuration"
	groupas "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/as/group"
	lifecyclehook "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/as/lifecyclehook"
	policy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/as/policy"
	server "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/bms/server"
	policycbr "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cbr/policy"
	vault "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cbr/vault"
	addon "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cce/addon"
	ccenamespace "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cce/ccenamespace"
	cluster "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cce/cluster"
	node "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cce/node"
	nodepool "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/cce/nodepool"
	alarmrule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ces/alarmrule"
	backup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/csbs/backup"
	backuppolicy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/csbs/backuppolicy"
	clustercss "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/css/cluster"
	snapshot "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/css/snapshot"
	instancedcs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dcs/instance"
	databaserole "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dds/databaserole"
	databaseuser "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dds/databaseuser"
	instancedds "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dds/instance"
	certificate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/certificate"
	ipgroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/ipgroup"
	listener "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/listener"
	loadbalancer "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/loadbalancer"
	member "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/member"
	monitor "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/monitor"
	pool "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dedicatedelb/pool"
	stream "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dis/stream"
	database "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/database"
	dlipackage "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/dlipackage"
	flinksqljob "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/flinksqljob"
	queue "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/queue"
	sparkjob "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/sparkjob"
	table "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dli/table"
	kafkainstance "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/kafkainstance"
	kafkatopic "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/kafkatopic"
	kafkauser "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/kafkauser"
	rocketmqconsumergroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/rocketmqconsumergroup"
	rocketmqinstance "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/rocketmqinstance"
	rocketmqtopic "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/rocketmqtopic"
	rocketmquser "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dms/rocketmquser"
	ptrrecord "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dns/ptrrecord"
	recordset "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dns/recordset"
	zone "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dns/zone"
	job "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/drs/job"
	clusterdws "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/dws/cluster"
	floatingipassociate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ecs/floatingipassociate"
	instanceecs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ecs/instance"
	interfaceattach "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ecs/interfaceattach"
	keypair "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ecs/keypair"
	servergroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ecs/servergroup"
	eip "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/eip/eip"
	eipassociate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/eip/eipassociate"
	l7policy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/l7policy"
	l7rule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/l7rule"
	listenerelb "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/listener"
	loadbalancerelb "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/loadbalancer"
	memberelb "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/member"
	monitorelb "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/monitor"
	poolelb "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/pool"
	whitelist "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/elb/whitelist"
	project "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/eps/project"
	blockstoragevolume "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/evs/blockstoragevolume"
	computevolumeattach "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/evs/computevolumeattach"
	function "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/fgs/function"
	trigger "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/fgs/trigger"
	acl "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/acl"
	agency "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/agency"
	groupiam "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/group"
	groupmembership "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/groupmembership"
	projectiam "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/project"
	provider "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/provider"
	providerconversion "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/providerconversion"
	role "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/role"
	roleassignment "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/roleassignment"
	user "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/iam/user"
	image "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/ims/image"
	key "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/kms/key"
	grouplts "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/lts/group"
	topic "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/lts/topic"
	vpcflowlog "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/lts/vpcflowlog"
	instancemls "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/mls/instance"
	dataset "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/modelarts/dataset"
	datasetversion "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/modelarts/datasetversion"
	clustermrs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/mrs/cluster"
	jobmrs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/mrs/job"
	dnatrule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/nat/dnatrule"
	gateway "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/nat/gateway"
	snatrule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/nat/snatrule"
	aclnetacl "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/netacl/acl"
	aclrule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/netacl/aclrule"
	obsbucket "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/obsbucket"
	obsbucketobject "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/obsbucketobject"
	obsbucketreplication "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/obsbucketreplication"
	s3bucket "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/s3bucket"
	s3bucketobject "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/s3bucketobject"
	s3bucketpolicy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/oss/s3bucketpolicy"
	providerconfig "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/providerconfig"
	account "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/account"
	databaserds "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/database"
	databaseprivilege "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/databaseprivilege"
	instancerds "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/instance"
	parametergroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/parametergroup"
	readreplica "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rds/readreplica"
	softwareconfig "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rts/softwareconfig"
	stack "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/rts/stack"
	drill "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sdrs/drill"
	protectedinstance "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sdrs/protectedinstance"
	protectiongroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sdrs/protectiongroup"
	replicationattach "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sdrs/replicationattach"
	replicationpair "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sdrs/replicationpair"
	accessrule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sfs/accessrule"
	filesystem "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sfs/filesystem"
	turbo "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sfs/turbo"
	subscription "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/smn/subscription"
	topicsmn "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/smn/topic"
	servertemplate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sms/servertemplate"
	task "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/sms/task"
	organization "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/swr/organization"
	organizationusers "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/swr/organizationusers"
	repository "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/swr/repository"
	repositorysharing "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/swr/repositorysharing"
	tags "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/tms/tags"
	backupvbs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vbs/backup"
	backuppolicyvbs "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vbs/backuppolicy"
	peeringconnection "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/peeringconnection"
	peeringconnectionaccepter "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/peeringconnectionaccepter"
	port "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/port"
	route "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/route"
	routetable "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/routetable"
	securitygroup "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/securitygrouprule"
	vip "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/vip"
	vipassociate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/vipassociate"
	vpc "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/vpc"
	vpcsubnet "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpc/vpcsubnet"
	approval "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpcep/approval"
	endpoint "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpcep/endpoint"
	vpcepservice "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/vpcep/vpcepservice"
	certificatewaf "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/certificate"
	dedicatedcertificate "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/dedicatedcertificate"
	dedicateddomain "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/dedicateddomain"
	dedicatedinstance "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/dedicatedinstance"
	dedicatedpolicy "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/dedicatedpolicy"
	domain "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/domain"
	policywaf "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/policy"
	rulealarmmasking "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/rulealarmmasking"
	ruleblacklist "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/ruleblacklist"
	ruleccprotection "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/ruleccprotection"
	ruledatamasking "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/ruledatamasking"
	rulepreciseprotection "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/rulepreciseprotection"
	rulewebtamperprotection "github.com/FlexibleEngineCloud/provider-flexibleengine/internal/controller/waf/rulewebtamperprotection"
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
		antiddos.Setup,
		configuration.Setup,
		groupas.Setup,
		lifecyclehook.Setup,
		policy.Setup,
		server.Setup,
		policycbr.Setup,
		vault.Setup,
		addon.Setup,
		ccenamespace.Setup,
		cluster.Setup,
		node.Setup,
		nodepool.Setup,
		alarmrule.Setup,
		backup.Setup,
		backuppolicy.Setup,
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
		stream.Setup,
		database.Setup,
		dlipackage.Setup,
		flinksqljob.Setup,
		queue.Setup,
		sparkjob.Setup,
		table.Setup,
		kafkainstance.Setup,
		kafkatopic.Setup,
		kafkauser.Setup,
		rocketmqconsumergroup.Setup,
		rocketmqinstance.Setup,
		rocketmqtopic.Setup,
		rocketmquser.Setup,
		ptrrecord.Setup,
		recordset.Setup,
		zone.Setup,
		job.Setup,
		clusterdws.Setup,
		floatingipassociate.Setup,
		instanceecs.Setup,
		interfaceattach.Setup,
		keypair.Setup,
		servergroup.Setup,
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
		computevolumeattach.Setup,
		function.Setup,
		trigger.Setup,
		acl.Setup,
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
		grouplts.Setup,
		topic.Setup,
		vpcflowlog.Setup,
		instancemls.Setup,
		dataset.Setup,
		datasetversion.Setup,
		clustermrs.Setup,
		jobmrs.Setup,
		dnatrule.Setup,
		gateway.Setup,
		snatrule.Setup,
		aclnetacl.Setup,
		aclrule.Setup,
		obsbucket.Setup,
		obsbucketobject.Setup,
		obsbucketreplication.Setup,
		s3bucket.Setup,
		s3bucketobject.Setup,
		s3bucketpolicy.Setup,
		providerconfig.Setup,
		account.Setup,
		databaserds.Setup,
		databaseprivilege.Setup,
		instancerds.Setup,
		parametergroup.Setup,
		readreplica.Setup,
		softwareconfig.Setup,
		stack.Setup,
		drill.Setup,
		protectedinstance.Setup,
		protectiongroup.Setup,
		replicationattach.Setup,
		replicationpair.Setup,
		accessrule.Setup,
		filesystem.Setup,
		turbo.Setup,
		subscription.Setup,
		topicsmn.Setup,
		servertemplate.Setup,
		task.Setup,
		organization.Setup,
		organizationusers.Setup,
		repository.Setup,
		repositorysharing.Setup,
		tags.Setup,
		backupvbs.Setup,
		backuppolicyvbs.Setup,
		peeringconnection.Setup,
		peeringconnectionaccepter.Setup,
		port.Setup,
		route.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		vip.Setup,
		vipassociate.Setup,
		vpc.Setup,
		vpcsubnet.Setup,
		approval.Setup,
		endpoint.Setup,
		vpcepservice.Setup,
		certificatewaf.Setup,
		dedicatedcertificate.Setup,
		dedicateddomain.Setup,
		dedicatedinstance.Setup,
		dedicatedpolicy.Setup,
		domain.Setup,
		policywaf.Setup,
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
