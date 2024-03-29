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

type NodeDataVolumesObservation struct {
}

type NodeDataVolumesParameters struct {

	// Specifies the disk expansion parameters in key/value pair format.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// Specifies the ID of a KMS key. This is used to encrypt the volume.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Specifies the disk size in GB.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the disk type.
	// +kubebuilder:validation:Required
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type NodeObservation struct {

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Private IP of the CCE node.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Public IP of the CCE node.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// ID of the ECS instance associated with the node.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Node status information.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NodeParameters struct {

	// Node annotation, key/value pair format. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// specify the name of the available partition (AZ).
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Bandwidth billing type. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Bandwidth size. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// +kubebuilder:validation:Optional
	BillingMode *float64 `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	// ID of the cluster. Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Represents the data disk to be created.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	DataVolumes []NodeDataVolumesParameters `json:"dataVolumes" tf:"data_volumes,omitempty"`

	// Number of elastic IPs to be dynamically created. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	EIPCount *float64 `json:"eipCount,omitempty" tf:"eip_count,omitempty"`

	// List of existing elastic IP IDs. Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	EIPIds []*string `json:"eipIds,omitempty" tf:"eip_ids,omitempty"`

	// References to EIP in eip to populate eipIds.
	// +kubebuilder:validation:Optional
	EIPIdsRefs []v1.Reference `json:"eipIdsRefs,omitempty" tf:"-"`

	// Selector for a list of EIP in eip to populate eipIds.
	// +kubebuilder:validation:Optional
	EIPIdsSelector *v1.Selector `json:"eipIdsSelector,omitempty" tf:"-"`

	// Specifies the ECS group ID. If specified, the node will be created under
	// the cloud server group. Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/ecs/v1beta1.ServerGroup
	// +kubebuilder:validation:Optional
	EcsGroupID *string `json:"ecsGroupId,omitempty" tf:"ecs_group_id,omitempty"`

	// Reference to a ServerGroup in ecs to populate ecsGroupId.
	// +kubebuilder:validation:Optional
	EcsGroupIDRef *v1.Reference `json:"ecsGroupIdRef,omitempty" tf:"-"`

	// Selector for a ServerGroup in ecs to populate ecsGroupId.
	// +kubebuilder:validation:Optional
	EcsGroupIDSelector *v1.Selector `json:"ecsGroupIdSelector,omitempty" tf:"-"`

	// Classification of cloud server specifications.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	EcsPerformanceType *string `json:"ecsPerformanceType,omitempty" tf:"ecs_performance_type,omitempty"`

	// Extended parameter. Changing this parameter will create a new resource.
	// Availiable keys:
	// +kubebuilder:validation:Optional
	ExtendParam map[string]*string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendParamChargingMode *float64 `json:"extendParamChargingMode,omitempty" tf:"extend_param_charging_mode,omitempty"`

	// Specifies the flavor id. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	FlavorID *string `json:"flavorId" tf:"flavor_id,omitempty"`

	// Elastic IP type.
	// +kubebuilder:validation:Optional
	Iptype *string `json:"iptype,omitempty" tf:"iptype,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/ecs/v1beta1.KeyPair
	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Reference to a KeyPair in ecs to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairRef *v1.Reference `json:"keyPairRef,omitempty" tf:"-"`

	// Selector for a KeyPair in ecs to populate keyPair.
	// +kubebuilder:validation:Optional
	KeyPairSelector *v1.Selector `json:"keyPairSelector,omitempty" tf:"-"`

	// Tags of a Kubernetes node, key/value pair format. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The maximum number of instances a node is allowed to create.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// Node Name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource ID in UUID format.
	// +kubebuilder:validation:Optional
	OrderID *string `json:"orderId,omitempty" tf:"order_id,omitempty"`

	// Operating System of the node, possible values are EulerOS 2.2 and CentOS 7.6. Defaults to EulerOS 2.2.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Script required after installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Postinstall *string `json:"postinstall,omitempty" tf:"postinstall,omitempty"`

	// Script required before installation. The input value can be a Base64 encoded string or not.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Preinstall *string `json:"preinstall,omitempty" tf:"preinstall,omitempty"`

	// The Product ID. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// The Public key. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// It corresponds to the system disk related configuration.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	RootVolume []NodeRootVolumeParameters `json:"rootVolume" tf:"root_volume,omitempty"`

	// Bandwidth sharing type. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Sharetype *string `json:"sharetype,omitempty" tf:"sharetype,omitempty"`

	// Specifies the ID of the VPC Subnet to which the NIC belongs.
	// Changing this parameter will create a new resource.
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

	// VM tag, key/value pair format.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// You can add taints to created nodes to configure anti-affinity.
	// Changing this parameter will create a new resource.
	// Each taint contains the following parameters:
	// +kubebuilder:validation:Optional
	Taints []NodeTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`
}

type NodeRootVolumeObservation struct {
}

type NodeRootVolumeParameters struct {

	// Specifies the disk expansion parameters in key/value pair format.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// Specifies the disk size in GB.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the disk type.
	// +kubebuilder:validation:Required
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type NodeTaintsObservation struct {
}

type NodeTaintsParameters struct {

	// Available options are NoSchedule, PreferNoSchedule, and NoExecute.
	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits,
	// hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// A value must start with a letter or digit and can contain a maximum of 63 characters,
	// including letters, digits, hyphens (-), underscores (_), and periods (.).
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeParameters `json:"forProvider"`
}

// NodeStatus defines the observed state of Node.
type NodeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Node is the Schema for the Nodes API. ""page_title: "flexibleengine_cce_node_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec   `json:"spec"`
	Status            NodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeList contains a list of Nodes
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

// Repository type metadata.
var (
	Node_Kind             = "Node"
	Node_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Node_Kind}.String()
	Node_KindAPIVersion   = Node_Kind + "." + CRDGroupVersion.String()
	Node_GroupVersionKind = CRDGroupVersion.WithKind(Node_Kind)
)

func init() {
	SchemeBuilder.Register(&Node{}, &NodeList{})
}
