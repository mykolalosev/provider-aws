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

type CertificateObservation struct {

	// The ARN of the certificate
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Set of domain validation objects which can be used to complete certificate validation. Can have more than one element, e.g., if SANs are defined. Only set if DNS-validation was used.
	DomainValidationOptions []DomainValidationOptionsObservation `json:"domainValidationOptions,omitempty" tf:"domain_validation_options,omitempty"`

	// The ARN of the certificate
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the certificate.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A list of addresses that received a validation E-Mail. Only set if EMAIL-validation was used.
	ValidationEmails []*string `json:"validationEmails,omitempty" tf:"validation_emails,omitempty"`
}

type CertificateParameters struct {

	// ARN of an ACM PCA
	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// The certificate's PEM-formatted public key
	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// The certificate's PEM-formatted chain
	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// A domain name for which the certificate should be issued
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Configuration block used to set certificate options. Detailed below.
	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// The certificate's PEM-formatted private key
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of domains that should be SANs in the issued certificate.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Which method to use for validation.
	// +kubebuilder:validation:Optional
	ValidationMethod *string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`

	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// +kubebuilder:validation:Optional
	ValidationOption []ValidationOptionParameters `json:"validationOption,omitempty" tf:"validation_option,omitempty"`
}

type DomainValidationOptionsObservation struct {

	// A fully qualified domain name (FQDN) in the certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The name of the DNS record to create to validate the certificate
	ResourceRecordName *string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`

	// The type of DNS record to create
	ResourceRecordType *string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`

	// The value the DNS record needs to have
	ResourceRecordValue *string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type DomainValidationOptionsParameters struct {
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// Specifies whether certificate details should be added to a certificate transparency log. Valid values are ENABLED or DISABLED. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
	// +kubebuilder:validation:Optional
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference,omitempty" tf:"certificate_transparency_logging_preference,omitempty"`
}

type ValidationOptionObservation struct {
}

type ValidationOptionParameters struct {

	// A fully qualified domain name (FQDN) in the certificate.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// The domain name that you want ACM to use to send you validation emails. This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the domain_name value or a superdomain of the domain_name value. For example, if you request a certificate for "testing.example.com", you can specify "example.com" for this value.
	// +kubebuilder:validation:Required
	ValidationDomain *string `json:"validationDomain" tf:"validation_domain,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Requests and manages a certificate from Amazon Certificate Manager (ACM).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
