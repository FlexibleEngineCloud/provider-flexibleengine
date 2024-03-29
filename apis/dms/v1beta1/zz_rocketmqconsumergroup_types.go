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

type RocketMQConsumerGroupObservation struct {

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RocketMQConsumerGroupParameters struct {

	// Specifies whether to broadcast of the consumer group.
	// Specifies whether to broadcast of the consumer group.
	// +kubebuilder:validation:Optional
	Broadcast *bool `json:"broadcast,omitempty" tf:"broadcast,omitempty"`

	// Specifies the list of associated brokers of the consumer group.
	// Specifies the list of associated brokers of the consumer group.
	// +kubebuilder:validation:Required
	Brokers []*string `json:"brokers" tf:"brokers,omitempty"`

	// Specifies the consumer group is enabled or not. Default to true.
	// Specifies the consumer group is enabled or not. Default to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the ID of the rocketMQ instance.
	// Specifies the ID of the rocketMQ instance.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/dms/v1beta1.RocketMQInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RocketMQInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RocketMQInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the consumer group.
	// Specifies the name of the consumer group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the maximum number of retry times.
	// Specifies the maximum number of retry times.
	// +kubebuilder:validation:Required
	RetryMaxTimes *float64 `json:"retryMaxTimes" tf:"retry_max_times,omitempty"`
}

// RocketMQConsumerGroupSpec defines the desired state of RocketMQConsumerGroup
type RocketMQConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RocketMQConsumerGroupParameters `json:"forProvider"`
}

// RocketMQConsumerGroupStatus defines the observed state of RocketMQConsumerGroup.
type RocketMQConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RocketMQConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RocketMQConsumerGroup is the Schema for the RocketMQConsumerGroups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RocketMQConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RocketMQConsumerGroupSpec   `json:"spec"`
	Status            RocketMQConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RocketMQConsumerGroupList contains a list of RocketMQConsumerGroups
type RocketMQConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RocketMQConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	RocketMQConsumerGroup_Kind             = "RocketMQConsumerGroup"
	RocketMQConsumerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RocketMQConsumerGroup_Kind}.String()
	RocketMQConsumerGroup_KindAPIVersion   = RocketMQConsumerGroup_Kind + "." + CRDGroupVersion.String()
	RocketMQConsumerGroup_GroupVersionKind = CRDGroupVersion.WithKind(RocketMQConsumerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RocketMQConsumerGroup{}, &RocketMQConsumerGroupList{})
}
