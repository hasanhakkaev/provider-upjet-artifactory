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

type RepositoryReplicationInitParameters struct {

	// Enabling the check_binary_existence_in_filestore flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see Optimizing Repository Replication with Checksum-Based Storage.
	// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `json:"checkBinaryExistenceInFilestore,omitempty" tf:"check_binary_existence_in_filestore,omitempty"`

	// A valid CRON expression that you can use to control replication frequency. Eg: 0 0 12 * * ? *, 0 0 2 ? * MON-SAT *. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year . Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by ?. Incorrect: * 5,7,9 14/2 * * WED,SAT *, correct: * 5,7,9 14/2 ? * WED,SAT *. See details in Cron Trigger Tutorial.
	// The Cron expression that determines when the next replication will be triggered.
	CronExp *string `json:"cronExp,omitempty" tf:"cron_exp,omitempty"`

	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is false.
	// com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication *bool `json:"enableEventReplication,omitempty" tf:"enable_event_replication,omitempty"`

	// When set, enables replication of this repository to the target specified in url attribute. Default value is true.
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default, no artifacts are excluded.
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludePathPrefixPattern *string `json:"excludePathPrefixPattern,omitempty" tf:"exclude_path_prefix_pattern,omitempty"`

	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludePathPrefixPattern *string `json:"includePathPrefixPattern,omitempty" tf:"include_path_prefix_pattern,omitempty"`

	// (Computed) Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	// Replication ID.
	ReplicationKey *string `json:"replicationKey,omitempty" tf:"replication_key,omitempty"`

	// Repository name.
	// Repository name.
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`

	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is false.
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes *bool `json:"syncDeletes,omitempty" tf:"sync_deletes,omitempty"`

	// When set, the task also synchronizes the properties of replicated artifacts. Default value is true.
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`
	SyncProperties *bool `json:"syncProperties,omitempty" tf:"sync_properties,omitempty"`
}

type RepositoryReplicationObservation struct {

	// Enabling the check_binary_existence_in_filestore flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see Optimizing Repository Replication with Checksum-Based Storage.
	// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `json:"checkBinaryExistenceInFilestore,omitempty" tf:"check_binary_existence_in_filestore,omitempty"`

	// A valid CRON expression that you can use to control replication frequency. Eg: 0 0 12 * * ? *, 0 0 2 ? * MON-SAT *. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year . Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by ?. Incorrect: * 5,7,9 14/2 * * WED,SAT *, correct: * 5,7,9 14/2 ? * WED,SAT *. See details in Cron Trigger Tutorial.
	// The Cron expression that determines when the next replication will be triggered.
	CronExp *string `json:"cronExp,omitempty" tf:"cron_exp,omitempty"`

	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is false.
	// com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication *bool `json:"enableEventReplication,omitempty" tf:"enable_event_replication,omitempty"`

	// When set, enables replication of this repository to the target specified in url attribute. Default value is true.
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default, no artifacts are excluded.
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludePathPrefixPattern *string `json:"excludePathPrefixPattern,omitempty" tf:"exclude_path_prefix_pattern,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludePathPrefixPattern *string `json:"includePathPrefixPattern,omitempty" tf:"include_path_prefix_pattern,omitempty"`

	// (Computed) Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	// Replication ID.
	ReplicationKey *string `json:"replicationKey,omitempty" tf:"replication_key,omitempty"`

	// Repository name.
	// Repository name.
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`

	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is false.
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes *bool `json:"syncDeletes,omitempty" tf:"sync_deletes,omitempty"`

	// When set, the task also synchronizes the properties of replicated artifacts. Default value is true.
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`
	SyncProperties *bool `json:"syncProperties,omitempty" tf:"sync_properties,omitempty"`
}

type RepositoryReplicationParameters struct {

	// Enabling the check_binary_existence_in_filestore flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see Optimizing Repository Replication with Checksum-Based Storage.
	// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	// +kubebuilder:validation:Optional
	CheckBinaryExistenceInFilestore *bool `json:"checkBinaryExistenceInFilestore,omitempty" tf:"check_binary_existence_in_filestore,omitempty"`

	// A valid CRON expression that you can use to control replication frequency. Eg: 0 0 12 * * ? *, 0 0 2 ? * MON-SAT *. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year . Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by ?. Incorrect: * 5,7,9 14/2 * * WED,SAT *, correct: * 5,7,9 14/2 ? * WED,SAT *. See details in Cron Trigger Tutorial.
	// The Cron expression that determines when the next replication will be triggered.
	// +kubebuilder:validation:Optional
	CronExp *string `json:"cronExp,omitempty" tf:"cron_exp,omitempty"`

	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is false.
	// com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	// +kubebuilder:validation:Optional
	EnableEventReplication *bool `json:"enableEventReplication,omitempty" tf:"enable_event_replication,omitempty"`

	// When set, enables replication of this repository to the target specified in url attribute. Default value is true.
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default, no artifacts are excluded.
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	// +kubebuilder:validation:Optional
	ExcludePathPrefixPattern *string `json:"excludePathPrefixPattern,omitempty" tf:"exclude_path_prefix_pattern,omitempty"`

	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	// +kubebuilder:validation:Optional
	IncludePathPrefixPattern *string `json:"includePathPrefixPattern,omitempty" tf:"include_path_prefix_pattern,omitempty"`

	// (Computed) Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	// Replication ID.
	// +kubebuilder:validation:Optional
	ReplicationKey *string `json:"replicationKey,omitempty" tf:"replication_key,omitempty"`

	// Repository name.
	// Repository name.
	// +kubebuilder:validation:Optional
	RepoKey *string `json:"repoKey,omitempty" tf:"repo_key,omitempty"`

	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is false.
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	// +kubebuilder:validation:Optional
	SyncDeletes *bool `json:"syncDeletes,omitempty" tf:"sync_deletes,omitempty"`

	// When set, the task also synchronizes the properties of replicated artifacts. Default value is true.
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`
	// +kubebuilder:validation:Optional
	SyncProperties *bool `json:"syncProperties,omitempty" tf:"sync_properties,omitempty"`
}

// RepositoryReplicationSpec defines the desired state of RepositoryReplication
type RepositoryReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryReplicationParameters `json:"forProvider"`
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
	InitProvider RepositoryReplicationInitParameters `json:"initProvider,omitempty"`
}

// RepositoryReplicationStatus defines the observed state of RepositoryReplication.
type RepositoryReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryReplication is the Schema for the RepositoryReplications API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,artifactory}
type RepositoryReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repoKey) || (has(self.initProvider) && has(self.initProvider.repoKey))",message="spec.forProvider.repoKey is a required parameter"
	Spec   RepositoryReplicationSpec   `json:"spec"`
	Status RepositoryReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryReplicationList contains a list of RepositoryReplications
type RepositoryReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryReplication `json:"items"`
}

// Repository type metadata.
var (
	RepositoryReplication_Kind             = "RepositoryReplication"
	RepositoryReplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryReplication_Kind}.String()
	RepositoryReplication_KindAPIVersion   = RepositoryReplication_Kind + "." + CRDGroupVersion.String()
	RepositoryReplication_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryReplication_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryReplication{}, &RepositoryReplicationList{})
}
