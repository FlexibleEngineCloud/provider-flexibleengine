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

type KeyPairObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KeyPairParameters struct {

	// +kubebuilder:validation:Optional
	PrivateKeyPath *string `json:"privateKeyPath,omitempty" tf:"private_key_path,omitempty"`

	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// KeyPairSpec defines the desired state of KeyPair
type KeyPairSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyPairParameters `json:"forProvider"`
}

// KeyPairStatus defines the observed state of KeyPair.
type KeyPairStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyPairObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPair is the Schema for the KeyPairs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type KeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyPairSpec   `json:"spec"`
	Status            KeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairList contains a list of KeyPairs
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPair `json:"items"`
}

// Repository type metadata.
var (
	KeyPair_Kind             = "KeyPair"
	KeyPair_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyPair_Kind}.String()
	KeyPair_KindAPIVersion   = KeyPair_Kind + "." + CRDGroupVersion.String()
	KeyPair_GroupVersionKind = CRDGroupVersion.WithKind(KeyPair_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyPair{}, &KeyPairList{})
}