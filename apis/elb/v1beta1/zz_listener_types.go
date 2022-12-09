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

type ListenerObservation struct {

	// The unique ID for the listener.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ListenerParameters struct {

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The ID of the default pool with which the
	// listener is associated. Changing this creates a new listener.
	// +crossplane:generate:reference:type=Pool
	// +kubebuilder:validation:Optional
	DefaultPoolID *string `json:"defaultPoolId,omitempty" tf:"default_pool_id,omitempty"`

	// Reference to a Pool to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDRef *v1.Reference `json:"defaultPoolIdRef,omitempty" tf:"-"`

	// Selector for a Pool to populate defaultPoolId.
	// +kubebuilder:validation:Optional
	DefaultPoolIDSelector *v1.Selector `json:"defaultPoolIdSelector,omitempty" tf:"-"`

	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is TERMINATED_HTTPS. See
	// here
	// for more information.
	// +kubebuilder:validation:Optional
	DefaultTLSContainerRef *string `json:"defaultTlsContainerRef,omitempty" tf:"default_tls_container_ref,omitempty"`

	// Human-readable description for the listener.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to use HTTP/2. The default value is false.
	// This parameter is valid only when the protocol is set to TERMINATED_HTTPS.
	// +kubebuilder:validation:Optional
	Http2Enable *bool `json:"http2Enable,omitempty" tf:"http2_enable,omitempty"`

	// Specifies the idle timeout duration, in seconds.
	// +kubebuilder:validation:Optional
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// The load balancer on which to provision this
	// listener. Changing this creates a new listener.
	// +crossplane:generate:reference:type=LoadBalancer
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// Human-readable name for the listener. Does not have
	// to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The protocol - can either be TCP, UDP, HTTP or TERMINATED_HTTPS.
	// Changing this creates a new listener.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The port on which to listen for client traffic.
	// Changing this creates a new listener.
	// +kubebuilder:validation:Required
	ProtocolPort *float64 `json:"protocolPort" tf:"protocol_port,omitempty"`

	// The region in which to create the listener resource.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new listener.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the timeout duration for waiting for a request from a client,
	// in seconds. This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300,
	// and the default value is 60.
	// +kubebuilder:validation:Optional
	RequestTimeout *float64 `json:"requestTimeout,omitempty" tf:"request_timeout,omitempty"`

	// Specifies the timeout duration for waiting for a request from a backend
	// server, in seconds. This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300,
	// and the default value is 60.
	// +kubebuilder:validation:Optional
	ResponseTimeout *float64 `json:"responseTimeout,omitempty" tf:"response_timeout,omitempty"`

	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// here
	// for more information.
	// +kubebuilder:validation:Optional
	SniContainerRefs []*string `json:"sniContainerRefs,omitempty" tf:"sni_container_refs,omitempty"`

	// Specifies the security policy used by the listener.
	// This parameter is valid only when the load balancer protocol is set to TERMINATED_HTTPS.
	// The value can be tls-1-0, tls-1-1, tls-1-2, or tls-1-2-strict, and the default value is tls-1-0.
	// For details of cipher suites for each security policy, see the table below.
	// +kubebuilder:validation:Optional
	TLSCiphersPolicy *string `json:"tlsCiphersPolicy,omitempty" tf:"tls_ciphers_policy,omitempty"`

	// The key/value pairs to associate with the listener.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique ID for the listener.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/iam/v1beta1.Project
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`

	// Specifies whether to pass source IP addresses of the clients to backend servers.
	// +kubebuilder:validation:Optional
	TransparentClientIPEnable *bool `json:"transparentClientIpEnable,omitempty" tf:"transparent_client_ip_enable,omitempty"`
}

// ListenerSpec defines the desired state of Listener
type ListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListenerParameters `json:"forProvider"`
}

// ListenerStatus defines the observed state of Listener.
type ListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Listener is the Schema for the Listeners API. ""page_title: "flexibleengine_lb_listener_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Listener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ListenerSpec   `json:"spec"`
	Status            ListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListenerList contains a list of Listeners
type ListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Listener `json:"items"`
}

// Repository type metadata.
var (
	Listener_Kind             = "Listener"
	Listener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Listener_Kind}.String()
	Listener_KindAPIVersion   = Listener_Kind + "." + CRDGroupVersion.String()
	Listener_GroupVersionKind = CRDGroupVersion.WithKind(Listener_Kind)
)

func init() {
	SchemeBuilder.Register(&Listener{}, &ListenerList{})
}
