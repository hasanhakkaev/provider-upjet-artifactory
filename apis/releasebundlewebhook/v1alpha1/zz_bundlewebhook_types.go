// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BundleWebhookInitParameters struct {

	// Specifies where the webhook will be applied, on which release bundles or distributions.
	Criteria []CriteriaInitParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Description of webhook. Max length 1000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Status of webhook. Default to 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook.
	// Allow values: created, signed, deleted
	// +listType=set
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	Handler []HandlerInitParameters `json:"handler,omitempty" tf:"handler,omitempty"`

	// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type BundleWebhookObservation struct {

	// Specifies where the webhook will be applied, on which release bundles or distributions.
	Criteria []CriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Description of webhook. Max length 1000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Status of webhook. Default to 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook.
	// Allow values: created, signed, deleted
	// +listType=set
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	Handler []HandlerObservation `json:"handler,omitempty" tf:"handler,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type BundleWebhookParameters struct {

	// Specifies where the webhook will be applied, on which release bundles or distributions.
	// +kubebuilder:validation:Optional
	Criteria []CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// Description of webhook. Max length 1000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Status of webhook. Default to 'true'
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook.
	// Allow values: created, signed, deleted
	// +kubebuilder:validation:Optional
	// +listType=set
	EventTypes []*string `json:"eventTypes,omitempty" tf:"event_types,omitempty"`

	// +kubebuilder:validation:Optional
	Handler []HandlerParameters `json:"handler,omitempty" tf:"handler,omitempty"`

	// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type CriteriaInitParameters struct {

	// Trigger on any release bundles or distributions
	AnyReleaseBundle *bool `json:"anyReleaseBundle,omitempty" tf:"any_release_bundle,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +listType=set
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +listType=set
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Trigger on this list of release bundle names
	// +listType=set
	RegisteredReleaseBundleNames []*string `json:"registeredReleaseBundleNames,omitempty" tf:"registered_release_bundle_names,omitempty"`
}

type CriteriaObservation struct {

	// Trigger on any release bundles or distributions
	AnyReleaseBundle *bool `json:"anyReleaseBundle,omitempty" tf:"any_release_bundle,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +listType=set
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +listType=set
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Trigger on this list of release bundle names
	// +listType=set
	RegisteredReleaseBundleNames []*string `json:"registeredReleaseBundleNames,omitempty" tf:"registered_release_bundle_names,omitempty"`
}

type CriteriaParameters struct {

	// Trigger on any release bundles or distributions
	// +kubebuilder:validation:Optional
	AnyReleaseBundle *bool `json:"anyReleaseBundle" tf:"any_release_bundle,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +kubebuilder:validation:Optional
	// +listType=set
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
	// +kubebuilder:validation:Optional
	// +listType=set
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Trigger on this list of release bundle names
	// +kubebuilder:validation:Optional
	// +listType=set
	RegisteredReleaseBundleNames []*string `json:"registeredReleaseBundleNames" tf:"registered_release_bundle_names,omitempty"`
}

type HandlerInitParameters struct {

	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	// +mapType=granular
	CustomHTTPHeaders map[string]*string `json:"customHttpHeaders,omitempty" tf:"custom_http_headers,omitempty"`

	// Proxy key from Artifactory UI (Administration -> Proxies -> Configuration)
	Proxy *string `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type HandlerObservation struct {

	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	// +mapType=granular
	CustomHTTPHeaders map[string]*string `json:"customHttpHeaders,omitempty" tf:"custom_http_headers,omitempty"`

	// Proxy key from Artifactory UI (Administration -> Proxies -> Configuration)
	Proxy *string `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type HandlerParameters struct {

	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CustomHTTPHeaders map[string]*string `json:"customHttpHeaders,omitempty" tf:"custom_http_headers,omitempty"`

	// Proxy key from Artifactory UI (Administration -> Proxies -> Configuration)
	// +kubebuilder:validation:Optional
	Proxy *string `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
	// +kubebuilder:validation:Optional
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// BundleWebhookSpec defines the desired state of BundleWebhook
type BundleWebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BundleWebhookParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BundleWebhookInitParameters `json:"initProvider,omitempty"`
}

// BundleWebhookStatus defines the observed state of BundleWebhook.
type BundleWebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BundleWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BundleWebhook is the Schema for the BundleWebhooks API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,artifactory}
type BundleWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.criteria) || (has(self.initProvider) && has(self.initProvider.criteria))",message="spec.forProvider.criteria is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventTypes) || (has(self.initProvider) && has(self.initProvider.eventTypes))",message="spec.forProvider.eventTypes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.handler) || (has(self.initProvider) && has(self.initProvider.handler))",message="spec.forProvider.handler is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   BundleWebhookSpec   `json:"spec"`
	Status BundleWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BundleWebhookList contains a list of BundleWebhooks
type BundleWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BundleWebhook `json:"items"`
}

// Repository type metadata.
var (
	BundleWebhook_Kind             = "BundleWebhook"
	BundleWebhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BundleWebhook_Kind}.String()
	BundleWebhook_KindAPIVersion   = BundleWebhook_Kind + "." + CRDGroupVersion.String()
	BundleWebhook_GroupVersionKind = CRDGroupVersion.WithKind(BundleWebhook_Kind)
)

func init() {
	SchemeBuilder.Register(&BundleWebhook{}, &BundleWebhookList{})
}
