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

type ApplicationObservation struct {

	// App key.
	// The APP key.
	AppKey *string `json:"appKey,omitempty" tf:"app_key,omitempty"`

	// ID of the APIG application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The registration time.
	RegistrationTime *string `json:"registrationTime,omitempty" tf:"registration_time,omitempty"`

	// The latest update time of the application.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ApplicationParameters struct {

	// Specifies an array of one or more application codes which the APIG application belongs
	// to. Up to five application codes can be created. The code consists of 64 to 180 characters, starting with a letter,
	// digit, plus sign (+) or slash (/). Only letters, digits and following special special characters are allowed: !@#$%+-_
	// /=
	// The array of one or more application codes that the application has.
	// +kubebuilder:validation:Optional
	AppCodes []*string `json:"appCodes,omitempty" tf:"app_codes,omitempty"`

	// Specifies the description about the APIG application. The description contain a
	// maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
	// Unicode format.
	// The application description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies an ID of the APIG dedicated instance to which the APIG
	// application belongs to. Changing this will create a new APIG application resource.
	// The ID of the dedicated instance to which the application belongs.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the API application. The API group name consists of 3 to 64
	// characters, starting with a letter. Only letters, digits and underscores (_) are allowed. Chinese characters must be
	// in UTF-8 or Unicode format.
	// The application name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the region in which to create the APIG application resource. If
	// omitted, the provider-level region will be used. Changing this will create a new APIG application resource.
	// The region where the application is located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the secret action to be done for the APIG application. The valid action
	// is RESET.
	// The secret action to be done for the application.
	// +kubebuilder:validation:Optional
	SecretAction *string `json:"secretAction,omitempty" tf:"secret_action,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. ""page_title: "flexibleengine_apig_application"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
