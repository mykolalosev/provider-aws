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

type ResourceShareObservation struct {

	// The Amazon Resource Name (ARN) of the resource share.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Amazon Resource Name (ARN) of the resource share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResourceShareParameters struct {

	// Indicates whether principals outside your organization can be associated with a resource share.
	// +kubebuilder:validation:Optional
	AllowExternalPrincipals *bool `json:"allowExternalPrincipals,omitempty" tf:"allow_external_principals,omitempty"`

	// The name of the resource share.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource share. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ResourceShareSpec defines the desired state of ResourceShare
type ResourceShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceShareParameters `json:"forProvider"`
}

// ResourceShareStatus defines the observed state of ResourceShare.
type ResourceShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceShare is the Schema for the ResourceShares API. Manages a Resource Access Manager (RAM) Resource Share.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceShareSpec   `json:"spec"`
	Status            ResourceShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceShareList contains a list of ResourceShares
type ResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceShare `json:"items"`
}

// Repository type metadata.
var (
	ResourceShare_Kind             = "ResourceShare"
	ResourceShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceShare_Kind}.String()
	ResourceShare_KindAPIVersion   = ResourceShare_Kind + "." + CRDGroupVersion.String()
	ResourceShare_GroupVersionKind = CRDGroupVersion.WithKind(ResourceShare_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceShare{}, &ResourceShareList{})
}
