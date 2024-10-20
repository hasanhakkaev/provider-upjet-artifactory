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

type OauthProviderInitParameters struct {

	// OAuth user info endpoint for the IdP.
	APIURL *string `json:"apiUrl,omitempty" tf:"api_url,omitempty"`

	// OAuth authorization endpoint for the IdP.
	AuthURL *string `json:"authUrl,omitempty" tf:"auth_url,omitempty"`

	// OAuth client ID configured on the IdP.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// OAuth client secret configured on the IdP.
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// Enable the Artifactory OAuth provider.  Default value is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name of the Artifactory OAuth provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OAuth token endpoint for the IdP.
	TokenURL *string `json:"tokenUrl,omitempty" tf:"token_url,omitempty"`

	// Type of OAuth provider. (e.g., github, google, cloudfoundry, or openId)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OauthProviderObservation struct {

	// OAuth user info endpoint for the IdP.
	APIURL *string `json:"apiUrl,omitempty" tf:"api_url,omitempty"`

	// OAuth authorization endpoint for the IdP.
	AuthURL *string `json:"authUrl,omitempty" tf:"auth_url,omitempty"`

	// OAuth client ID configured on the IdP.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Enable the Artifactory OAuth provider.  Default value is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name of the Artifactory OAuth provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OAuth token endpoint for the IdP.
	TokenURL *string `json:"tokenUrl,omitempty" tf:"token_url,omitempty"`

	// Type of OAuth provider. (e.g., github, google, cloudfoundry, or openId)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OauthProviderParameters struct {

	// OAuth user info endpoint for the IdP.
	// +kubebuilder:validation:Optional
	APIURL *string `json:"apiUrl" tf:"api_url,omitempty"`

	// OAuth authorization endpoint for the IdP.
	// +kubebuilder:validation:Optional
	AuthURL *string `json:"authUrl" tf:"auth_url,omitempty"`

	// OAuth client ID configured on the IdP.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// OAuth client secret configured on the IdP.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// Enable the Artifactory OAuth provider.  Default value is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name of the Artifactory OAuth provider.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// OAuth token endpoint for the IdP.
	// +kubebuilder:validation:Optional
	TokenURL *string `json:"tokenUrl" tf:"token_url,omitempty"`

	// Type of OAuth provider. (e.g., github, google, cloudfoundry, or openId)
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SettingsInitParameters struct {

	// Allow persisted users to access their profile.  Default value is false.
	AllowUserToAccessProfile *bool `json:"allowUserToAccessProfile,omitempty" tf:"allow_user_to_access_profile,omitempty"`

	// Enable OAuth SSO.  Default value is true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProvider []OauthProviderInitParameters `json:"oauthProvider,omitempty" tf:"oauth_provider,omitempty"`

	// Enable the creation of local Artifactory users.  Default value is false.
	PersistUsers *bool `json:"persistUsers,omitempty" tf:"persist_users,omitempty"`
}

type SettingsObservation struct {

	// Allow persisted users to access their profile.  Default value is false.
	AllowUserToAccessProfile *bool `json:"allowUserToAccessProfile,omitempty" tf:"allow_user_to_access_profile,omitempty"`

	// Enable OAuth SSO.  Default value is true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProvider []OauthProviderObservation `json:"oauthProvider,omitempty" tf:"oauth_provider,omitempty"`

	// Enable the creation of local Artifactory users.  Default value is false.
	PersistUsers *bool `json:"persistUsers,omitempty" tf:"persist_users,omitempty"`
}

type SettingsParameters struct {

	// Allow persisted users to access their profile.  Default value is false.
	// +kubebuilder:validation:Optional
	AllowUserToAccessProfile *bool `json:"allowUserToAccessProfile,omitempty" tf:"allow_user_to_access_profile,omitempty"`

	// Enable OAuth SSO.  Default value is true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	// +kubebuilder:validation:Optional
	OauthProvider []OauthProviderParameters `json:"oauthProvider,omitempty" tf:"oauth_provider,omitempty"`

	// Enable the creation of local Artifactory users.  Default value is false.
	// +kubebuilder:validation:Optional
	PersistUsers *bool `json:"persistUsers,omitempty" tf:"persist_users,omitempty"`
}

// SettingsSpec defines the desired state of Settings
type SettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SettingsParameters `json:"forProvider"`
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
	InitProvider SettingsInitParameters `json:"initProvider,omitempty"`
}

// SettingsStatus defines the observed state of Settings.
type SettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Settings is the Schema for the Settingss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,artifactory}
type Settings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.oauthProvider) || (has(self.initProvider) && has(self.initProvider.oauthProvider))",message="spec.forProvider.oauthProvider is a required parameter"
	Spec   SettingsSpec   `json:"spec"`
	Status SettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SettingsList contains a list of Settingss
type SettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Settings `json:"items"`
}

// Repository type metadata.
var (
	Settings_Kind             = "Settings"
	Settings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Settings_Kind}.String()
	Settings_KindAPIVersion   = Settings_Kind + "." + CRDGroupVersion.String()
	Settings_GroupVersionKind = CRDGroupVersion.WithKind(Settings_Kind)
)

func init() {
	SchemeBuilder.Register(&Settings{}, &SettingsList{})
}
