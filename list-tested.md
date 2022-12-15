
## csbs
|Kind|CRD|Tested|
|---|---|---|
|Backup|Backup.csbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|BackupPolicy|BackupPolicy.csbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## oss
|Kind|CRD|Tested|
|---|---|---|
|S3BucketObject|S3BucketObject.oss.flexibleengine.upbound.io/v1beta1|:x:|
|OBSBucketObject|OBSBucketObject.oss.flexibleengine.upbound.io/v1beta1|:x:|
|OBSBucket|OBSBucket.oss.flexibleengine.upbound.io/v1beta1|:x:|
|S3Bucket|S3Bucket.oss.flexibleengine.upbound.io/v1beta1|:x:|
|OBSBucketReplication|OBSBucketReplication.oss.flexibleengine.upbound.io/v1beta1|:x:|
|S3BucketPolicy|S3BucketPolicy.oss.flexibleengine.upbound.io/v1beta1|:x:|

## vpc
|Kind|CRD|Tested|
|---|---|---|
|PeeringConnectionAccepter|PeeringConnectionAccepter.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|NetworkingSubnet|NetworkingSubnet.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Router|Router.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VIP|VIP.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Port|Port.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RouteTable|RouteTable.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|Network|Network.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VPC|VPC.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|SecurityGroupRule|SecurityGroupRule.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|FlowLog|FlowLog.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|RouterInterface|RouterInterface.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|PeeringConnection|PeeringConnection.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VPCSubnet|VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VIPAssociate|VIPAssociate.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Route|Route.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|SecurityGroup|SecurityGroup.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## sfs
|Kind|CRD|Tested|
|---|---|---|
|Turbo|Turbo.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|AccessRule|AccessRule.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|FileSystem|FileSystem.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## rds
|Kind|CRD|Tested|
|---|---|---|
|ReadReplica|ReadReplica.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Account|Account.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Database|Database.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Instance|Instance.rds.flexibleengine.upbound.io/v1beta1|:x:|
|ParameterGroup|ParameterGroup.rds.flexibleengine.upbound.io/v1beta1|:x:|
|DatabasePrivilege|DatabasePrivilege.rds.flexibleengine.upbound.io/v1beta1|:x:|

## nat
|Kind|CRD|Tested|
|---|---|---|
|SnatRule|SnatRule.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Gateway|Gateway.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DnatRule|DnatRule.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## eip
|Kind|CRD|Tested|
|---|---|---|
|EIPAssociate|EIPAssociate.eip.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|EIP|EIP.eip.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## waf
|Kind|CRD|Tested|
|---|---|---|
|Policy|Policy.waf.flexibleengine.upbound.io/v1beta1|:x:|
|Domain|Domain.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RuleDataMasking|RuleDataMasking.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RulePreciseProtection|RulePreciseProtection.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RuleAlarmMasking|RuleAlarmMasking.waf.flexibleengine.upbound.io/v1beta1|:x:|
|Certificate|Certificate.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RuleWebTamperProtection|RuleWebTamperProtection.waf.flexibleengine.upbound.io/v1beta1|:x:|
|DedicatedInstance|DedicatedInstance.waf.flexibleengine.upbound.io/v1beta1|:x:|
|DedicatedPolicy|DedicatedPolicy.waf.flexibleengine.upbound.io/v1beta1|:x:|
|DedicatedDomain|DedicatedDomain.waf.flexibleengine.upbound.io/v1beta1|:x:|
|DedicatedCertificate|DedicatedCertificate.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RuleBlacklist|RuleBlacklist.waf.flexibleengine.upbound.io/v1beta1|:x:|
|RuleCcProtection|RuleCcProtection.waf.flexibleengine.upbound.io/v1beta1|:x:|

## vpcep
|Kind|CRD|Tested|
|---|---|---|
|Endpoint|Endpoint.vpcep.flexibleengine.upbound.io/v1beta1|:x:|
|Approval|Approval.vpcep.flexibleengine.upbound.io/v1beta1|:x:|
|Service|Service.vpcep.flexibleengine.upbound.io/v1beta1|:x:|

## dli
|Kind|CRD|Tested|
|---|---|---|
|Table|Table.dli.flexibleengine.upbound.io/v1beta1|:x:|
|FlinksqlJob|FlinksqlJob.dli.flexibleengine.upbound.io/v1beta1|:x:|
|Database|Database.dli.flexibleengine.upbound.io/v1beta1|:x:|
|SparkJob|SparkJob.dli.flexibleengine.upbound.io/v1beta1|:x:|
|DLIPackage|DLIPackage.dli.flexibleengine.upbound.io/v1beta1|:x:|
|Queue|Queue.dli.flexibleengine.upbound.io/v1beta1|:x:|

## flexibleengine
|Kind|CRD|Tested|
|---|---|---|
|StoreConfig|StoreConfig.flexibleengine.upbound.io/v1alpha1|:x:|
|ProviderConfig|ProviderConfig.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ProviderConfigUsage|ProviderConfigUsage.flexibleengine.upbound.io/v1beta1|:x:|

## fgs
|Kind|CRD|Tested|
|---|---|---|
|Trigger|Trigger.fgs.flexibleengine.upbound.io/v1beta1|:x:|
|Dependency|Dependency.fgs.flexibleengine.upbound.io/v1beta1|:x:|
|Function|Function.fgs.flexibleengine.upbound.io/v1beta1|:x:|

## dns
|Kind|CRD|Tested|
|---|---|---|
|Zone|Zone.dns.flexibleengine.upbound.io/v1beta1|:x:|
|Recordset|Recordset.dns.flexibleengine.upbound.io/v1beta1|:x:|

## ecs
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|FloatingIpAssociate|FloatingIpAssociate.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|KeyPair|KeyPair.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ServerGroup|ServerGroup.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|InterfaceAttach|InterfaceAttach.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|FloatingIp|FloatingIp.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|VolumeAttach|VolumeAttach.ecs.flexibleengine.upbound.io/v1beta1|:x:|

## iam
|Kind|CRD|Tested|
|---|---|---|
|Agency|Agency.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ProviderConversion|ProviderConversion.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Project|Project.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Group|Group.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RoleAssignment|RoleAssignment.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|User|User.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|GroupMembership|GroupMembership.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Provider|Provider.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Role|Role.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## mls
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.mls.flexibleengine.upbound.io/v1beta1|:x:|

## eps
|Kind|CRD|Tested|
|---|---|---|
|Project|Project.eps.flexibleengine.upbound.io/v1beta1|:x:|

## swr
|Kind|CRD|Tested|
|---|---|---|
|RepositorySharing|RepositorySharing.swr.flexibleengine.upbound.io/v1beta1|:x:|
|Repository|Repository.swr.flexibleengine.upbound.io/v1beta1|:x:|
|Organization|Organization.swr.flexibleengine.upbound.io/v1beta1|:x:|
|OrganizationUsers|OrganizationUsers.swr.flexibleengine.upbound.io/v1beta1|:x:|

## dcs
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.dcs.flexibleengine.upbound.io/v1beta1|:x:|

## agd
|Kind|CRD|Tested|
|---|---|---|
|Group|Group.agd.flexibleengine.upbound.io/v1beta1|:x:|
|APIPublishment|APIPublishment.agd.flexibleengine.upbound.io/v1beta1|:x:|
|VPCChannel|VPCChannel.agd.flexibleengine.upbound.io/v1beta1|:x:|
|API|API.agd.flexibleengine.upbound.io/v1beta1|:x:|
|Instance|Instance.agd.flexibleengine.upbound.io/v1beta1|:x:|
|Response|Response.agd.flexibleengine.upbound.io/v1beta1|:x:|
|CustomAuthorizer|CustomAuthorizer.agd.flexibleengine.upbound.io/v1beta1|:x:|
|Application|Application.agd.flexibleengine.upbound.io/v1beta1|:x:|
|Environment|Environment.agd.flexibleengine.upbound.io/v1beta1|:x:|
|ThrottlingPolicy|ThrottlingPolicy.agd.flexibleengine.upbound.io/v1beta1|:x:|
|ThrottlingPolicyAssociate|ThrottlingPolicyAssociate.agd.flexibleengine.upbound.io/v1beta1|:x:|

## dis
|Kind|CRD|Tested|
|---|---|---|
|Stream|Stream.dis.flexibleengine.upbound.io/v1beta1|:x:|

## netacl
|Kind|CRD|Tested|
|---|---|---|
|FirewallGroup|FirewallGroup.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|ACLRule|ACLRule.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|Policy|Policy.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|ACL|ACL.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|Rule|Rule.netacl.flexibleengine.upbound.io/v1beta1|:x:|

## sdrs
|Kind|CRD|Tested|
|---|---|---|
|Drill|Drill.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ReplicationPair|ReplicationPair.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ProtectionGroup|ProtectionGroup.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ReplicationAttach|ReplicationAttach.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ProtectedInstance|ProtectedInstance.sdrs.flexibleengine.upbound.io/v1beta1|:x:|

## mrs
|Kind|CRD|Tested|
|---|---|---|
|Cluster|Cluster.mrs.flexibleengine.upbound.io/v1beta1|:x:|
|Job|Job.mrs.flexibleengine.upbound.io/v1beta1|:x:|

## dds
|Kind|CRD|Tested|
|---|---|---|
|DatabaseRole|DatabaseRole.dds.flexibleengine.upbound.io/v1beta1|:x:|
|Instance|Instance.dds.flexibleengine.upbound.io/v1beta1|:x:|
|DatabaseUser|DatabaseUser.dds.flexibleengine.upbound.io/v1beta1|:x:|

## ag
|Kind|CRD|Tested|
|---|---|---|
|Group|Group.ag.flexibleengine.upbound.io/v1beta1|:x:|
|API|API.ag.flexibleengine.upbound.io/v1beta1|:x:|

## rts
|Kind|CRD|Tested|
|---|---|---|
|Stack|Stack.rts.flexibleengine.upbound.io/v1beta1|:x:|
|SoftwareConfig|SoftwareConfig.rts.flexibleengine.upbound.io/v1beta1|:x:|

## elb
|Kind|CRD|Tested|
|---|---|---|
|L7Rule|L7Rule.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|L7Policy|L7Policy.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|LoadBalancer|LoadBalancer.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Member|Member.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Whitelist|Whitelist.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Listener|Listener.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Monitor|Monitor.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Pool|Pool.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## ims
|Kind|CRD|Tested|
|---|---|---|
|Image|Image.ims.flexibleengine.upbound.io/v1beta1|:x:|

## cse
|Kind|CRD|Tested|
|---|---|---|
|Microservice|Microservice.cse.flexibleengine.upbound.io/v1beta1|:x:|
|MicroserviceEngine|MicroserviceEngine.cse.flexibleengine.upbound.io/v1beta1|:x:|
|MicroserviceInstance|MicroserviceInstance.cse.flexibleengine.upbound.io/v1beta1|:x:|

## cce
|Kind|CRD|Tested|
|---|---|---|
|Node|Node.cce.flexibleengine.upbound.io/v1beta1|:x:|
|NodePool|NodePool.cce.flexibleengine.upbound.io/v1beta1|:x:|
|Cluster|Cluster.cce.flexibleengine.upbound.io/v1beta1|:x:|
|Pvc|Pvc.cce.flexibleengine.upbound.io/v1beta1|:x:|
|Addon|Addon.cce.flexibleengine.upbound.io/v1beta1|:x:|
|Namespace|Namespace.cce.flexibleengine.upbound.io/v1beta1|:x:|

## smn
|Kind|CRD|Tested|
|---|---|---|
|Topic|Topic.smn.flexibleengine.upbound.io/v1beta1|:x:|
|Subscription|Subscription.smn.flexibleengine.upbound.io/v1beta1|:x:|

## dms
|Kind|CRD|Tested|
|---|---|---|
|KafkaUser|KafkaUser.dms.flexibleengine.upbound.io/v1beta1|:x:|
|KafkaTopic|KafkaTopic.dms.flexibleengine.upbound.io/v1beta1|:x:|
|KafkaInstance|KafkaInstance.dms.flexibleengine.upbound.io/v1beta1|:x:|

## as
|Kind|CRD|Tested|
|---|---|---|
|LifecycleHook|LifecycleHook.as.flexibleengine.upbound.io/v1beta1|:x:|
|Policy|Policy.as.flexibleengine.upbound.io/v1beta1|:x:|
|Group|Group.as.flexibleengine.upbound.io/v1beta1|:x:|
|Configuration|Configuration.as.flexibleengine.upbound.io/v1beta1|:x:|

## modelarts
|Kind|CRD|Tested|
|---|---|---|
|DatasetVersion|DatasetVersion.modelarts.flexibleengine.upbound.io/v1beta1|:x:|
|Dataset|Dataset.modelarts.flexibleengine.upbound.io/v1beta1|:x:|

## kms
|Kind|CRD|Tested|
|---|---|---|
|Key|Key.kms.flexibleengine.upbound.io/v1beta1|:x:|

## vbs
|Kind|CRD|Tested|
|---|---|---|
|BackupPolicy|BackupPolicy.vbs.flexibleengine.upbound.io/v1beta1|:x:|
|Backup|Backup.vbs.flexibleengine.upbound.io/v1beta1|:x:|

## dws
|Kind|CRD|Tested|
|---|---|---|
|Cluster|Cluster.dws.flexibleengine.upbound.io/v1beta1|:x:|

## dedicatedelb
|Kind|CRD|Tested|
|---|---|---|
|Member|Member.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|Certificate|Certificate.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|Pool|Pool.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|LoadBalancer|LoadBalancer.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|Monitor|Monitor.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|Listener|Listener.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|
|IPGroup|IPGroup.dedicatedelb.flexibleengine.upbound.io/v1beta1|:x:|

## lts
|Kind|CRD|Tested|
|---|---|---|
|Group|Group.lts.flexibleengine.upbound.io/v1beta1|:x:|
|Topic|Topic.lts.flexibleengine.upbound.io/v1beta1|:x:|

## cbr
|Kind|CRD|Tested|
|---|---|---|
|Vault|Vault.cbr.flexibleengine.upbound.io/v1beta1|:x:|
|Policy|Policy.cbr.flexibleengine.upbound.io/v1beta1|:x:|

## css
|Kind|CRD|Tested|
|---|---|---|
|Snapshot|Snapshot.css.flexibleengine.upbound.io/v1beta1|:x:|
|Cluster|Cluster.css.flexibleengine.upbound.io/v1beta1|:x:|

## bms
|Kind|CRD|Tested|
|---|---|---|
|Server|Server.bms.flexibleengine.upbound.io/v1beta1|:x:|

## tms
|Kind|CRD|Tested|
|---|---|---|
|Tags|Tags.tms.flexibleengine.upbound.io/v1beta1|:x:|

## mrsd
|Kind|CRD|Tested|
|---|---|---|
|HybridCluster|HybridCluster.mrsd.flexibleengine.upbound.io/v1beta1|:x:|
|Cluster|Cluster.mrsd.flexibleengine.upbound.io/v1beta1|:x:|
|Job|Job.mrsd.flexibleengine.upbound.io/v1beta1|:x:|

## antiddos
|Kind|CRD|Tested|
|---|---|---|
|AntiDDoS|AntiDDoS.antiddos.flexibleengine.upbound.io/v1beta1|:x:|

## evs
|Kind|CRD|Tested|
|---|---|---|
|BlockStorageVolume|BlockStorageVolume.evs.flexibleengine.upbound.io/v1beta1|:x:|

## cts
|Kind|CRD|Tested|
|---|---|---|
|Tracker|Tracker.cts.flexibleengine.upbound.io/v1beta1|:x:|

## drs
|Kind|CRD|Tested|
|---|---|---|
|Job|Job.drs.flexibleengine.upbound.io/v1beta1|:x:|

## ces
|Kind|CRD|Tested|
|---|---|---|
|AlarmRule|AlarmRule.ces.flexibleengine.upbound.io/v1beta1|:x:|
