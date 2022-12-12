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

type ImageObservation struct {

	// The checksum of the data associated with the image.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// The date the image was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// A unique ID assigned by Glance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The id of the flexibleengine user who owns the image.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The path to the JSON-schema that represent
	// the image or image
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The size in bytes of the data associated with the image.
	SizeBytes *float64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`

	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The date the image was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ImageParameters struct {

	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	// +kubebuilder:validation:Required
	ContainerFormat *string `json:"containerFormat" tf:"container_format,omitempty"`

	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	// +kubebuilder:validation:Required
	DiskFormat *string `json:"diskFormat" tf:"disk_format,omitempty"`

	// This is the directory where the images will
	// be downloaded. Images will be stored with a filename corresponding to
	// the url's md5 hash. Defaults to "$HOME/
	// +kubebuilder:validation:Optional
	ImageCachePath *string `json:"imageCachePath,omitempty" tf:"image_cache_path,omitempty"`

	// This is the url of the raw image that will
	// be downloaded in the image_cache_path before being uploaded to Glance.
	// Glance is able to download image from internet but the gophercloud library
	// does not yet provide a way to do so.
	// Conflicts with local_file_path.
	// +kubebuilder:validation:Optional
	ImageSourceURL *string `json:"imageSourceUrl,omitempty" tf:"image_source_url,omitempty"`

	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with image_source_url.
	// +kubebuilder:validation:Optional
	LocalFilePath *string `json:"localFilePath,omitempty" tf:"local_file_path,omitempty"`

	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	// +kubebuilder:validation:Optional
	MinDiskGb *float64 `json:"minDiskGb,omitempty" tf:"min_disk_gb,omitempty"`

	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	// +kubebuilder:validation:Optional
	MinRAMMb *float64 `json:"minRamMb,omitempty" tf:"min_ram_mb,omitempty"`

	// The name of the image.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// If true, image will not be deletable.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the region argument of the provider
	// is used. Changing this creates a new Image.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the FlexibleEngine cloud.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. ""page_title: "flexibleengine_images_image_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}