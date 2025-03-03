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

type AuthenticateCognitoObservation struct {
}

type AuthenticateCognitoParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Behavior if the user is not authenticated. Valid values are deny, allow and authenticate.
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// ARN of the Cognito user pool.
	// +kubebuilder:validation:Required
	UserPoolArn *string `json:"userPoolArn" tf:"user_pool_arn,omitempty"`

	// ID of the Cognito user pool client.
	// +kubebuilder:validation:Required
	UserPoolClientID *string `json:"userPoolClientId" tf:"user_pool_client_id,omitempty"`

	// Domain prefix or fully-qualified domain name of the Cognito user pool.
	// +kubebuilder:validation:Required
	UserPoolDomain *string `json:"userPoolDomain" tf:"user_pool_domain,omitempty"`
}

type AuthenticateOidcObservation struct {
}

type AuthenticateOidcParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Authorization endpoint of the IdP.
	// +kubebuilder:validation:Required
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// OAuth 2.0 client identifier.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// OAuth 2.0 client secret.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// OIDC issuer identifier of the IdP.
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// Behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// Token endpoint of the IdP.
	// +kubebuilder:validation:Required
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// User info endpoint of the IdP.
	// +kubebuilder:validation:Required
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint,omitempty"`
}

type DefaultActionObservation struct {
}

type DefaultActionParameters struct {

	// Configuration block for using Amazon Cognito to authenticate users. Specify only when type is authenticate-cognito. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticateCognito []AuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when type is authenticate-oidc. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticateOidc []AuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	// +kubebuilder:validation:Optional
	FixedResponse []FixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn. Detailed below.
	// +kubebuilder:validation:Optional
	Forward []ForwardParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between 1 and 50000.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Configuration block for creating a redirect action. Required if type is redirect. Detailed below.
	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// ARN of the Target Group to which to route traffic. Specify only if type is forward and you want to route to a single target group. To route to one or more target groups, use a forward block instead.
	// +crossplane:generate:reference:type=LBTargetGroup
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// Reference to a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnRef *v1.Reference `json:"targetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnSelector *v1.Selector `json:"targetGroupArnSelector,omitempty" tf:"-"`

	// Type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FixedResponseObservation struct {
}

type FixedResponseParameters struct {

	// Content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Message body.
	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// HTTP response code. Valid values are 2XX, 4XX, or 5XX.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ForwardObservation struct {
}

type ForwardParameters struct {

	// Configuration block for target group stickiness for the rule. Detailed below.
	// +kubebuilder:validation:Optional
	Stickiness []StickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Set of 1-5 target group blocks. Detailed below.
	// +kubebuilder:validation:Required
	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group,omitempty"`
}

type LBListenerObservation struct {

	// ARN of the listener (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the listener (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBListenerParameters struct {

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if protocol is TLS. Valid values are HTTP1Only, HTTP2Only, HTTP2Optional, HTTP2Preferred, and None.
	// +kubebuilder:validation:Optional
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy,omitempty"`

	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the aws_lb_listener_certificate resource.
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Configuration block for default actions. Detailed below.
	// +kubebuilder:validation:Required
	DefaultAction []DefaultActionParameters `json:"defaultAction" tf:"default_action,omitempty"`

	// ARN of the load balancer.
	// +crossplane:generate:reference:type=LB
	// +kubebuilder:validation:Optional
	LoadBalancerArn *string `json:"loadBalancerArn,omitempty" tf:"load_balancer_arn,omitempty"`

	// Reference to a LB to populate loadBalancerArn.
	// +kubebuilder:validation:Optional
	LoadBalancerArnRef *v1.Reference `json:"loadBalancerArnRef,omitempty" tf:"-"`

	// Selector for a LB to populate loadBalancerArn.
	// +kubebuilder:validation:Optional
	LoadBalancerArnSelector *v1.Selector `json:"loadBalancerArnSelector,omitempty" tf:"-"`

	// Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are HTTP and HTTPS, with a default of HTTP. For Network Load Balancers, valid values are TCP, TLS, UDP, and TCP_UDP. Not valid to use UDP or TCP_UDP if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the SSL Policy for the listener. Required if protocol is HTTPS or TLS.
	// +kubebuilder:validation:Optional
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {

	// Hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type StickinessObservation struct {
}

type StickinessParameters struct {

	// Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	// +kubebuilder:validation:Required
	Duration *float64 `json:"duration" tf:"duration,omitempty"`

	// Whether target group stickiness is enabled. Default is false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TargetGroupObservation struct {
}

type TargetGroupParameters struct {

	// ARN of the target group.
	// +crossplane:generate:reference:type=LBTargetGroup
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a LBTargetGroup to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Weight. The range is 0 to 999.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// LBListenerSpec defines the desired state of LBListener
type LBListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBListenerParameters `json:"forProvider"`
}

// LBListenerStatus defines the observed state of LBListener.
type LBListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBListener is the Schema for the LBListeners API. Provides a Load Balancer Listener resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBListenerSpec   `json:"spec"`
	Status            LBListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerList contains a list of LBListeners
type LBListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBListener `json:"items"`
}

// Repository type metadata.
var (
	LBListener_Kind             = "LBListener"
	LBListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBListener_Kind}.String()
	LBListener_KindAPIVersion   = LBListener_Kind + "." + CRDGroupVersion.String()
	LBListener_GroupVersionKind = CRDGroupVersion.WithKind(LBListener_Kind)
)

func init() {
	SchemeBuilder.Register(&LBListener{}, &LBListenerList{})
}
