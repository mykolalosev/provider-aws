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

type FargateProfileObservation struct {

	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// EKS Cluster name and EKS Fargate Profile name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the EKS Fargate Profile.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FargateProfileParameters struct {

	// 100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (^[0-9A-Za-z][A-Za-z0-9\-_]+$).
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// –  Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PodExecutionRoleArn *string `json:"podExecutionRoleArn,omitempty" tf:"pod_execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate podExecutionRoleArn.
	// +kubebuilder:validation:Optional
	PodExecutionRoleArnRef *v1.Reference `json:"podExecutionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate podExecutionRoleArn.
	// +kubebuilder:validation:Optional
	PodExecutionRoleArnSelector *v1.Selector `json:"podExecutionRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	// +kubebuilder:validation:Required
	Selector []SelectorParameters `json:"selector" tf:"selector,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: kubernetes.io/cluster/CLUSTER_NAME (where CLUSTER_NAME is replaced with the name of the EKS Cluster).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SelectorObservation struct {
}

type SelectorParameters struct {

	// Key-value map of Kubernetes labels for selection.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Kubernetes namespace for selection.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

// FargateProfileSpec defines the desired state of FargateProfile
type FargateProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FargateProfileParameters `json:"forProvider"`
}

// FargateProfileStatus defines the observed state of FargateProfile.
type FargateProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FargateProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FargateProfile is the Schema for the FargateProfiles API. Manages an EKS Fargate Profile
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargateProfileSpec   `json:"spec"`
	Status            FargateProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FargateProfileList contains a list of FargateProfiles
type FargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FargateProfile `json:"items"`
}

// Repository type metadata.
var (
	FargateProfile_Kind             = "FargateProfile"
	FargateProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FargateProfile_Kind}.String()
	FargateProfile_KindAPIVersion   = FargateProfile_Kind + "." + CRDGroupVersion.String()
	FargateProfile_GroupVersionKind = CRDGroupVersion.WithKind(FargateProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&FargateProfile{}, &FargateProfileList{})
}
