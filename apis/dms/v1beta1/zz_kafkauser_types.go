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

type KafkaUserObservation struct {

	// The resource ID which is formatted <instance_id>/<user_name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KafkaUserParameters struct {

	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ecs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the user. Changing this creates a new resource.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// KafkaUserSpec defines the desired state of KafkaUser
type KafkaUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaUserParameters `json:"forProvider"`
}

// KafkaUserStatus defines the observed state of KafkaUser.
type KafkaUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaUser is the Schema for the KafkaUsers API. ""page_title: "flexibleengine_dms_kafka_user"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type KafkaUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KafkaUserSpec   `json:"spec"`
	Status            KafkaUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaUserList contains a list of KafkaUsers
type KafkaUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaUser `json:"items"`
}

// Repository type metadata.
var (
	KafkaUser_Kind             = "KafkaUser"
	KafkaUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaUser_Kind}.String()
	KafkaUser_KindAPIVersion   = KafkaUser_Kind + "." + CRDGroupVersion.String()
	KafkaUser_GroupVersionKind = CRDGroupVersion.WithKind(KafkaUser_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaUser{}, &KafkaUserList{})
}