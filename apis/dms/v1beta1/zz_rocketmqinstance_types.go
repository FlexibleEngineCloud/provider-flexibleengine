/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CrossVPCAccessesObservation struct {
	AdvertisedIP *string `json:"advertisedIp,omitempty" tf:"advertised_ip,omitempty"`

	LisenterIP *string `json:"lisenterIp,omitempty" tf:"lisenter_ip,omitempty"`

	ListenerIP *string `json:"listenerIp,omitempty" tf:"listener_ip,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies a resource ID in UUID format.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`
}

type CrossVPCAccessesParameters struct {
}

type RocketMQInstanceObservation struct {

	// Indicates the service data address.
	// Indicates the service data address.
	BrokerAddress *string `json:"brokerAddress,omitempty" tf:"broker_address,omitempty"`

	CrossVPCAccesses []CrossVPCAccessesObservation `json:"crossVpcAccesses,omitempty" tf:"cross_vpc_accesses,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the time at which the maintenance window starts. The format is HH:mm:ss.
	// Indicates the time at which the maintenance window starts. The format is HH:mm:ss.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Indicates the time at which the maintenance window ends. The format is HH:mm:ss.
	// Indicates the time at which the maintenance window ends. The format is HH:mm:ss.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Indicates the metadata address.
	// Indicates the metadata address.
	NamesrvAddress *string `json:"namesrvAddress,omitempty" tf:"namesrv_address,omitempty"`

	// Indicates whether billing based on new specifications is enabled.
	// Indicates whether billing based on new specifications is enabled.
	NewSpecBillingEnable *bool `json:"newSpecBillingEnable,omitempty" tf:"new_spec_billing_enable,omitempty"`

	// Indicates the node quantity.
	// Indicates the node quantity.
	NodeNum *float64 `json:"nodeNum,omitempty" tf:"node_num,omitempty"`

	// Indicates the public network service data address.
	// Indicates the public network service data address.
	PublicBrokerAddress *string `json:"publicBrokerAddress,omitempty" tf:"public_broker_address,omitempty"`

	// Indicates the public network metadata address.
	// Indicates the public network metadata address.
	PublicNamesrvAddress *string `json:"publicNamesrvAddress,omitempty" tf:"public_namesrv_address,omitempty"`

	// Indicates the public IP address.
	// Indicates the public IP address.
	PublicipAddress *string `json:"publicipAddress,omitempty" tf:"publicip_address,omitempty"`

	// Indicates the resource specifications.
	// Indicates the resource specifications.
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty" tf:"resource_spec_code,omitempty"`

	// Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications
	// and the number of nodes are returned.
	// Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications
	// and the number of nodes are returned.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the status of the DMS RocketMQ instance.
	// Indicates the status of the DMS RocketMQ instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the DMS RocketMQ instance type. Value: cluster.
	// Indicates the DMS RocketMQ instance type. Value: cluster.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the used message storage space. Unit: GB.
	// Indicates the used message storage space. Unit: GB.
	UsedStorageSpace *float64 `json:"usedStorageSpace,omitempty" tf:"used_storage_space,omitempty"`
}

type RocketMQInstanceParameters struct {

	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the list of availability zone names, where
	// instance brokers reside and which has available resources.
	// Specifies the list of availability zone names
	// +kubebuilder:validation:Required
	AvailabilityZones []*string `json:"availabilityZones" tf:"availability_zones,omitempty"`

	// Specifies the broker numbers. Defaults to 1.
	// Changing this parameter will create a new resource.
	// Specifies the broker numbers.
	// +kubebuilder:validation:Optional
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the DMS RocketMQ instance.
	// The description can contain a maximum of 1024 characters.
	// Specifies the description of the DMS RocketMQ instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether access control is enabled.
	// Specifies whether access control is enabled.
	// +kubebuilder:validation:Optional
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies whether to enable public access.
	// By default, public access is disabled.
	// Changing this parameter will create a new resource.
	// Specifies whether to enable public access.
	// +kubebuilder:validation:Optional
	EnablePublicip *bool `json:"enablePublicip,omitempty" tf:"enable_publicip,omitempty"`

	// Specifies the version of the RocketMQ engine. Value: 4.8.0.
	// Changing this parameter will create a new resource.
	// Specifies the version of the RocketMQ engine.
	// +kubebuilder:validation:Required
	EngineVersion *string `json:"engineVersion" tf:"engine_version,omitempty"`

	// Specifies a resource ID in UUID format.
	// Specifies the enterprise project id of the instance.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies a product ID. The options are as follows:
	// Specifies a product ID
	// +kubebuilder:validation:Required
	FlavorID *string `json:"flavorId" tf:"flavor_id,omitempty"`

	// Specifies whether to support IPv6. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether to support IPv6
	// +kubebuilder:validation:Optional
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// Specifies the name of the DMS RocketMQ instance.
	// An instance name starts with a letter, consists of 4 to 64 characters, and can contain only letters,
	// digits, underscores (_), and hyphens (-).
	// Specifies the name of the DMS RocketMQ instance
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies the ID of the EIP bound to the instance.
	// Use commas (,) to separate multiple EIP IDs.
	// This parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).
	// Changing this parameter will create a new resource.
	// Specifies the ID of the EIP bound to the instance.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	PublicipID *string `json:"publicipId,omitempty" tf:"publicip_id,omitempty"`

	// Reference to a EIP in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDRef *v1.Reference `json:"publicipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate publicipId.
	// +kubebuilder:validation:Optional
	PublicipIDSelector *v1.Selector `json:"publicipIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the ACL access control.
	// Specifies whether access control is enabled.
	// +kubebuilder:validation:Optional
	RetentionPolicy *bool `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Specifies whether the RocketMQ SASL_SSL is enabled. Defaults to false.
	// Changing this parameter will create a new resource.
	// Specifies whether the RocketMQ SASL_SSL is enabled.
	// +kubebuilder:validation:Optional
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// Specifies the ID of a security group
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRef
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the message storage capacity, Unit: GB.
	// Value range: 300-3000.
	// Changing this parameter will create a new resource.
	// Specifies the message storage capacity, Unit: GB.
	// +kubebuilder:validation:Required
	StorageSpace *float64 `json:"storageSpace" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// The options are as follows:
	// Specifies the storage I/O specification
	// +kubebuilder:validation:Required
	StorageSpecCode *string `json:"storageSpecCode" tf:"storage_spec_code,omitempty"`

	// Specifies the ID of a subnet.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a subnet
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +crossplane:generate:reference:refFieldName=SubnetIDRef
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of a VPC.
	// Changing this parameter will create a new resource.
	// Specifies the ID of a VPC
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RocketMQInstanceSpec defines the desired state of RocketMQInstance
type RocketMQInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RocketMQInstanceParameters `json:"forProvider"`
}

// RocketMQInstanceStatus defines the observed state of RocketMQInstance.
type RocketMQInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RocketMQInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RocketMQInstance is the Schema for the RocketMQInstances API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RocketMQInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RocketMQInstanceSpec   `json:"spec"`
	Status            RocketMQInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RocketMQInstanceList contains a list of RocketMQInstances
type RocketMQInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RocketMQInstance `json:"items"`
}

// Repository type metadata.
var (
	RocketMQInstance_Kind             = "RocketMQInstance"
	RocketMQInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RocketMQInstance_Kind}.String()
	RocketMQInstance_KindAPIVersion   = RocketMQInstance_Kind + "." + CRDGroupVersion.String()
	RocketMQInstance_GroupVersionKind = CRDGroupVersion.WithKind(RocketMQInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RocketMQInstance{}, &RocketMQInstanceList{})
}
