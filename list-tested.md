# Resources tested

Last update: Ven  6 jan 2023 08:01:07 CET

## vpc
|Kind|CRD|Tested|
|---|---|---|
|SecurityGroup|SecurityGroup.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Route|Route.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|VPCSubnet|VPCSubnet.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|PeeringConnection|PeeringConnection.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|FlowLog|FlowLog.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|VPC|VPC.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VIP|VIP.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Port|Port.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VIPAssociate|VIPAssociate.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RouteTable|RouteTable.vpc.flexibleengine.upbound.io/v1beta1|:x:|
|SecurityGroupRule|SecurityGroupRule.vpc.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|PeeringConnectionAccepter|PeeringConnectionAccepter.vpc.flexibleengine.upbound.io/v1beta1|:x:|

## waf
|Kind|CRD|Tested|
|---|---|---|
|Certificate|Certificate.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RuleDataMasking|RuleDataMasking.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RuleBlacklist|RuleBlacklist.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DedicatedPolicy|DedicatedPolicy.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RuleCcProtection|RuleCcProtection.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Policy|Policy.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RulePreciseProtection|RulePreciseProtection.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DedicatedDomain|DedicatedDomain.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DedicatedCertificate|DedicatedCertificate.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DedicatedInstance|DedicatedInstance.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Domain|Domain.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RuleWebTamperProtection|RuleWebTamperProtection.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RuleAlarmMasking|RuleAlarmMasking.waf.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## ecs
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|InterfaceAttach|InterfaceAttach.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|KeyPair|KeyPair.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VolumeAttach|VolumeAttach.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|FloatingIpAssociate|FloatingIpAssociate.ecs.flexibleengine.upbound.io/v1beta1|:x:|
|ServerGroup|ServerGroup.ecs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## cts
|Kind|CRD|Tested|
|---|---|---|
|Tracker|Tracker.cts.flexibleengine.upbound.io/v1beta1|:x:|

## mrs
|Kind|CRD|Tested|
|---|---|---|
|Cluster|Cluster.mrs.flexibleengine.upbound.io/v1beta1|:x:|
|Job|Job.mrs.flexibleengine.upbound.io/v1beta1|:x:|

## netacl
|Kind|CRD|Tested|
|---|---|---|
|ACLRule|ACLRule.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|ACL|ACL.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|FirewallGroup|FirewallGroup.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|Policy|Policy.netacl.flexibleengine.upbound.io/v1beta1|:x:|
|Rule|Rule.netacl.flexibleengine.upbound.io/v1beta1|:x:|

## sdrs
|Kind|CRD|Tested|
|---|---|---|
|ReplicationAttach|ReplicationAttach.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ReplicationPair|ReplicationPair.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|Drill|Drill.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ProtectionGroup|ProtectionGroup.sdrs.flexibleengine.upbound.io/v1beta1|:x:|
|ProtectedInstance|ProtectedInstance.sdrs.flexibleengine.upbound.io/v1beta1|:x:|

## drs
|Kind|CRD|Tested|
|---|---|---|
|Job|Job.drs.flexibleengine.upbound.io/v1beta1|:x:|

## modelarts
|Kind|CRD|Tested|
|---|---|---|
|Dataset|Dataset.modelarts.flexibleengine.upbound.io/v1beta1|:x:|
|DatasetVersion|DatasetVersion.modelarts.flexibleengine.upbound.io/v1beta1|:x:|

## dms
|Kind|CRD|Tested|
|---|---|---|
|KafkaTopic|KafkaTopic.dms.flexibleengine.upbound.io/v1beta1|:x:|
|KafkaInstance|KafkaInstance.dms.flexibleengine.upbound.io/v1beta1|:x:|
|KafkaUser|KafkaUser.dms.flexibleengine.upbound.io/v1beta1|:x:|

## cbr
|Kind|CRD|Tested|
|---|---|---|
|Vault|Vault.cbr.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Policy|Policy.cbr.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## nat
|Kind|CRD|Tested|
|---|---|---|
|Gateway|Gateway.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DnatRule|DnatRule.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|SnatRule|SnatRule.nat.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## mrsd
|Kind|CRD|Tested|
|---|---|---|
|Job|Job.mrsd.flexibleengine.upbound.io/v1beta1|:x:|
|HybridCluster|HybridCluster.mrsd.flexibleengine.upbound.io/v1beta1|:x:|
|Cluster|Cluster.mrsd.flexibleengine.upbound.io/v1beta1|:x:|

## kms
|Kind|CRD|Tested|
|---|---|---|
|Key|Key.kms.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## dli
|Kind|CRD|Tested|
|---|---|---|
|SparkJob|SparkJob.dli.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Database|Database.dli.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Table|Table.dli.flexibleengine.upbound.io/v1beta1|:x:|
|Queue|Queue.dli.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|FlinksqlJob|FlinksqlJob.dli.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|DLIPackage|DLIPackage.dli.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## dds
|Kind|CRD|Tested|
|---|---|---|
|DatabaseUser|DatabaseUser.dds.flexibleengine.upbound.io/v1beta1|:x:|
|DatabaseRole|DatabaseRole.dds.flexibleengine.upbound.io/v1beta1|:x:|
|Instance|Instance.dds.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## tms
|Kind|CRD|Tested|
|---|---|---|
|Tags|Tags.tms.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## flexibleengine
|Kind|CRD|Tested|
|---|---|---|
|ProviderConfig|ProviderConfig.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ProviderConfigUsage|ProviderConfigUsage.flexibleengine.upbound.io/v1beta1|:x:|
|StoreConfig|StoreConfig.flexibleengine.upbound.io/v1alpha1|:x:|

## elb
|Kind|CRD|Tested|
|---|---|---|
|Listener|Listener.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Whitelist|Whitelist.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|L7Rule|L7Rule.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|LoadBalancer|LoadBalancer.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Pool|Pool.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Monitor|Monitor.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Member|Member.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|L7Policy|L7Policy.elb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## dedicatedelb
|Kind|CRD|Tested|
|---|---|---|
|Member|Member.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Pool|Pool.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Listener|Listener.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Certificate|Certificate.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Monitor|Monitor.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|LoadBalancer|LoadBalancer.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|IPGroup|IPGroup.dedicatedelb.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## css
|Kind|CRD|Tested|
|---|---|---|
|Snapshot|Snapshot.css.flexibleengine.upbound.io/v1beta1|:x:|
|Cluster|Cluster.css.flexibleengine.upbound.io/v1beta1|:x:|

## agd
|Kind|CRD|Tested|
|---|---|---|
|ThrottlingPolicyAssociate|ThrottlingPolicyAssociate.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|API|API.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Instance|Instance.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Group|Group.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|APIPublishment|APIPublishment.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Environment|Environment.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|VPCChannel|VPCChannel.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Application|Application.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ThrottlingPolicy|ThrottlingPolicy.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|CustomAuthorizer|CustomAuthorizer.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Response|Response.agd.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## iam
|Kind|CRD|Tested|
|---|---|---|
|Project|Project.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Agency|Agency.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Role|Role.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|GroupMembership|GroupMembership.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|RoleAssignment|RoleAssignment.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Group|Group.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Provider|Provider.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|User|User.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|ProviderConversion|ProviderConversion.iam.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## rds
|Kind|CRD|Tested|
|---|---|---|
|DatabasePrivilege|DatabasePrivilege.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Account|Account.rds.flexibleengine.upbound.io/v1beta1|:x:|
|ReadReplica|ReadReplica.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Database|Database.rds.flexibleengine.upbound.io/v1beta1|:x:|
|Instance|Instance.rds.flexibleengine.upbound.io/v1beta1|:x:|
|ParameterGroup|ParameterGroup.rds.flexibleengine.upbound.io/v1beta1|:x:|

## eip
|Kind|CRD|Tested|
|---|---|---|
|EIPAssociate|EIPAssociate.eip.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|EIP|EIP.eip.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## as
|Kind|CRD|Tested|
|---|---|---|
|Group|Group.as.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|LifecycleHook|LifecycleHook.as.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Policy|Policy.as.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Configuration|Configuration.as.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## ces
|Kind|CRD|Tested|
|---|---|---|
|AlarmRule|AlarmRule.ces.flexibleengine.upbound.io/v1beta1|:x:|

## sfs
|Kind|CRD|Tested|
|---|---|---|
|FileSystem|FileSystem.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Turbo|Turbo.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|AccessRule|AccessRule.sfs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## cce
|Kind|CRD|Tested|
|---|---|---|
|Addon|Addon.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|NodePool|NodePool.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Node|Node.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Cluster|Cluster.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Namespace|Namespace.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Pvc|Pvc.cce.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## mls
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.mls.flexibleengine.upbound.io/v1beta1|:x:|

## lts
|Kind|CRD|Tested|
|---|---|---|
|Group|Group.lts.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Topic|Topic.lts.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## vpcep
|Kind|CRD|Tested|
|---|---|---|
|VPCEPService|VPCEPService.vpcep.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Approval|Approval.vpcep.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Endpoint|Endpoint.vpcep.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## smn
|Kind|CRD|Tested|
|---|---|---|
|Subscription|Subscription.smn.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Topic|Topic.smn.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## oss
|Kind|CRD|Tested|
|---|---|---|
|S3BucketPolicy|S3BucketPolicy.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|S3BucketObject|S3BucketObject.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|OBSBucket|OBSBucket.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|OBSBucketObject|OBSBucketObject.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|S3Bucket|S3Bucket.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|OBSBucketReplication|OBSBucketReplication.oss.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## dws
|Kind|CRD|Tested|
|---|---|---|
|Cluster|Cluster.dws.flexibleengine.upbound.io/v1beta1|:x:|

## antiddos
|Kind|CRD|Tested|
|---|---|---|
|AntiDDoS|AntiDDoS.antiddos.flexibleengine.upbound.io/v1beta1|:x:|

## fgs
|Kind|CRD|Tested|
|---|---|---|
|Trigger|Trigger.fgs.flexibleengine.upbound.io/v1beta1|:x:|
|Function|Function.fgs.flexibleengine.upbound.io/v1beta1|:x:|
|Dependency|Dependency.fgs.flexibleengine.upbound.io/v1beta1|:x:|

## rts
|Kind|CRD|Tested|
|---|---|---|
|SoftwareConfig|SoftwareConfig.rts.flexibleengine.upbound.io/v1beta1|:x:|
|Stack|Stack.rts.flexibleengine.upbound.io/v1beta1|:x:|

## bms
|Kind|CRD|Tested|
|---|---|---|
|Server|Server.bms.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## swr
|Kind|CRD|Tested|
|---|---|---|
|OrganizationUsers|OrganizationUsers.swr.flexibleengine.upbound.io/v1beta1|:x:|
|RepositorySharing|RepositorySharing.swr.flexibleengine.upbound.io/v1beta1|:x:|
|Organization|Organization.swr.flexibleengine.upbound.io/v1beta1|:x:|
|Repository|Repository.swr.flexibleengine.upbound.io/v1beta1|:x:|

## dns
|Kind|CRD|Tested|
|---|---|---|
|Recordset|Recordset.dns.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Zone|Zone.dns.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Ptrrecord|Ptrrecord.dns.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## ag
|Kind|CRD|Tested|
|---|---|---|
|API|API.ag.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Group|Group.ag.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## vbs
|Kind|CRD|Tested|
|---|---|---|
|BackupPolicy|BackupPolicy.vbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|Backup|Backup.vbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## ims
|Kind|CRD|Tested|
|---|---|---|
|Image|Image.ims.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## cse
|Kind|CRD|Tested|
|---|---|---|
|MicroserviceEngine|MicroserviceEngine.cse.flexibleengine.upbound.io/v1beta1|:x:|
|MicroserviceInstance|MicroserviceInstance.cse.flexibleengine.upbound.io/v1beta1|:x:|
|Microservice|Microservice.cse.flexibleengine.upbound.io/v1beta1|:x:|

## csbs
|Kind|CRD|Tested|
|---|---|---|
|Backup|Backup.csbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
|BackupPolicy|BackupPolicy.csbs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## eps
|Kind|CRD|Tested|
|---|---|---|
|Project|Project.eps.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## evs
|Kind|CRD|Tested|
|---|---|---|
|BlockStorageVolume|BlockStorageVolume.evs.flexibleengine.upbound.io/v1beta1|:white_check_mark:|

## dcs
|Kind|CRD|Tested|
|---|---|---|
|Instance|Instance.dcs.flexibleengine.upbound.io/v1beta1|:x:|

## dis
|Kind|CRD|Tested|
|---|---|---|
|Stream|Stream.dis.flexibleengine.upbound.io/v1beta1|:white_check_mark:|
