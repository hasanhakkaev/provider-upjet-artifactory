//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepository) DeepCopyInto(out *ChefRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepository.
func (in *ChefRepository) DeepCopy() *ChefRepository {
	if in == nil {
		return nil
	}
	out := new(ChefRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChefRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositoryInitParameters) DeepCopyInto(out *ChefRepositoryInitParameters) {
	*out = *in
	if in.AllowAnyHostAuth != nil {
		in, out := &in.AllowAnyHostAuth, &out.AllowAnyHostAuth
		*out = new(bool)
		**out = **in
	}
	if in.AssumedOfflinePeriodSecs != nil {
		in, out := &in.AssumedOfflinePeriodSecs, &out.AssumedOfflinePeriodSecs
		*out = new(float64)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.BlockMismatchingMimeTypes != nil {
		in, out := &in.BlockMismatchingMimeTypes, &out.BlockMismatchingMimeTypes
		*out = new(bool)
		**out = **in
	}
	if in.BypassHeadRequests != nil {
		in, out := &in.BypassHeadRequests, &out.BypassHeadRequests
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.ClientTLSCertificate != nil {
		in, out := &in.ClientTLSCertificate, &out.ClientTLSCertificate
		*out = new(string)
		**out = **in
	}
	if in.ContentSynchronisation != nil {
		in, out := &in.ContentSynchronisation, &out.ContentSynchronisation
		*out = make([]ContentSynchronisationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableProxy != nil {
		in, out := &in.DisableProxy, &out.DisableProxy
		*out = new(bool)
		**out = **in
	}
	if in.DisableURLNormalization != nil {
		in, out := &in.DisableURLNormalization, &out.DisableURLNormalization
		*out = new(bool)
		**out = **in
	}
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.EnableCookieManagement != nil {
		in, out := &in.EnableCookieManagement, &out.EnableCookieManagement
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
		*out = new(string)
		**out = **in
	}
	if in.HardFail != nil {
		in, out := &in.HardFail, &out.HardFail
		*out = new(bool)
		**out = **in
	}
	if in.IncludesPattern != nil {
		in, out := &in.IncludesPattern, &out.IncludesPattern
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ListRemoteFolderItems != nil {
		in, out := &in.ListRemoteFolderItems, &out.ListRemoteFolderItems
		*out = new(bool)
		**out = **in
	}
	if in.LocalAddress != nil {
		in, out := &in.LocalAddress, &out.LocalAddress
		*out = new(string)
		**out = **in
	}
	if in.MetadataRetrievalTimeoutSecs != nil {
		in, out := &in.MetadataRetrievalTimeoutSecs, &out.MetadataRetrievalTimeoutSecs
		*out = new(float64)
		**out = **in
	}
	if in.MismatchingMimeTypesOverrideList != nil {
		in, out := &in.MismatchingMimeTypesOverrideList, &out.MismatchingMimeTypesOverrideList
		*out = new(string)
		**out = **in
	}
	if in.MissedCachePeriodSeconds != nil {
		in, out := &in.MissedCachePeriodSeconds, &out.MissedCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.Offline != nil {
		in, out := &in.Offline, &out.Offline
		*out = new(bool)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PriorityResolution != nil {
		in, out := &in.PriorityResolution, &out.PriorityResolution
		*out = new(bool)
		**out = **in
	}
	if in.ProjectEnvironments != nil {
		in, out := &in.ProjectEnvironments, &out.ProjectEnvironments
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.PropertySets != nil {
		in, out := &in.PropertySets, &out.PropertySets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = new(string)
		**out = **in
	}
	if in.RemoteRepoLayoutRef != nil {
		in, out := &in.RemoteRepoLayoutRef, &out.RemoteRepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RetrievalCachePeriodSeconds != nil {
		in, out := &in.RetrievalCachePeriodSeconds, &out.RetrievalCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ShareConfiguration != nil {
		in, out := &in.ShareConfiguration, &out.ShareConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.SocketTimeoutMillis != nil {
		in, out := &in.SocketTimeoutMillis, &out.SocketTimeoutMillis
		*out = new(float64)
		**out = **in
	}
	if in.StoreArtifactsLocally != nil {
		in, out := &in.StoreArtifactsLocally, &out.StoreArtifactsLocally
		*out = new(bool)
		**out = **in
	}
	if in.SynchronizeProperties != nil {
		in, out := &in.SynchronizeProperties, &out.SynchronizeProperties
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UnusedArtifactsCleanupPeriodHours != nil {
		in, out := &in.UnusedArtifactsCleanupPeriodHours, &out.UnusedArtifactsCleanupPeriodHours
		*out = new(float64)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositoryInitParameters.
func (in *ChefRepositoryInitParameters) DeepCopy() *ChefRepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(ChefRepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositoryList) DeepCopyInto(out *ChefRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChefRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositoryList.
func (in *ChefRepositoryList) DeepCopy() *ChefRepositoryList {
	if in == nil {
		return nil
	}
	out := new(ChefRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChefRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositoryObservation) DeepCopyInto(out *ChefRepositoryObservation) {
	*out = *in
	if in.AllowAnyHostAuth != nil {
		in, out := &in.AllowAnyHostAuth, &out.AllowAnyHostAuth
		*out = new(bool)
		**out = **in
	}
	if in.AssumedOfflinePeriodSecs != nil {
		in, out := &in.AssumedOfflinePeriodSecs, &out.AssumedOfflinePeriodSecs
		*out = new(float64)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.BlockMismatchingMimeTypes != nil {
		in, out := &in.BlockMismatchingMimeTypes, &out.BlockMismatchingMimeTypes
		*out = new(bool)
		**out = **in
	}
	if in.BypassHeadRequests != nil {
		in, out := &in.BypassHeadRequests, &out.BypassHeadRequests
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.ClientTLSCertificate != nil {
		in, out := &in.ClientTLSCertificate, &out.ClientTLSCertificate
		*out = new(string)
		**out = **in
	}
	if in.ContentSynchronisation != nil {
		in, out := &in.ContentSynchronisation, &out.ContentSynchronisation
		*out = make([]ContentSynchronisationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableProxy != nil {
		in, out := &in.DisableProxy, &out.DisableProxy
		*out = new(bool)
		**out = **in
	}
	if in.DisableURLNormalization != nil {
		in, out := &in.DisableURLNormalization, &out.DisableURLNormalization
		*out = new(bool)
		**out = **in
	}
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.EnableCookieManagement != nil {
		in, out := &in.EnableCookieManagement, &out.EnableCookieManagement
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
		*out = new(string)
		**out = **in
	}
	if in.HardFail != nil {
		in, out := &in.HardFail, &out.HardFail
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IncludesPattern != nil {
		in, out := &in.IncludesPattern, &out.IncludesPattern
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ListRemoteFolderItems != nil {
		in, out := &in.ListRemoteFolderItems, &out.ListRemoteFolderItems
		*out = new(bool)
		**out = **in
	}
	if in.LocalAddress != nil {
		in, out := &in.LocalAddress, &out.LocalAddress
		*out = new(string)
		**out = **in
	}
	if in.MetadataRetrievalTimeoutSecs != nil {
		in, out := &in.MetadataRetrievalTimeoutSecs, &out.MetadataRetrievalTimeoutSecs
		*out = new(float64)
		**out = **in
	}
	if in.MismatchingMimeTypesOverrideList != nil {
		in, out := &in.MismatchingMimeTypesOverrideList, &out.MismatchingMimeTypesOverrideList
		*out = new(string)
		**out = **in
	}
	if in.MissedCachePeriodSeconds != nil {
		in, out := &in.MissedCachePeriodSeconds, &out.MissedCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.Offline != nil {
		in, out := &in.Offline, &out.Offline
		*out = new(bool)
		**out = **in
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.PriorityResolution != nil {
		in, out := &in.PriorityResolution, &out.PriorityResolution
		*out = new(bool)
		**out = **in
	}
	if in.ProjectEnvironments != nil {
		in, out := &in.ProjectEnvironments, &out.ProjectEnvironments
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.PropertySets != nil {
		in, out := &in.PropertySets, &out.PropertySets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = new(string)
		**out = **in
	}
	if in.RemoteRepoLayoutRef != nil {
		in, out := &in.RemoteRepoLayoutRef, &out.RemoteRepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RetrievalCachePeriodSeconds != nil {
		in, out := &in.RetrievalCachePeriodSeconds, &out.RetrievalCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ShareConfiguration != nil {
		in, out := &in.ShareConfiguration, &out.ShareConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.SocketTimeoutMillis != nil {
		in, out := &in.SocketTimeoutMillis, &out.SocketTimeoutMillis
		*out = new(float64)
		**out = **in
	}
	if in.StoreArtifactsLocally != nil {
		in, out := &in.StoreArtifactsLocally, &out.StoreArtifactsLocally
		*out = new(bool)
		**out = **in
	}
	if in.SynchronizeProperties != nil {
		in, out := &in.SynchronizeProperties, &out.SynchronizeProperties
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UnusedArtifactsCleanupPeriodHours != nil {
		in, out := &in.UnusedArtifactsCleanupPeriodHours, &out.UnusedArtifactsCleanupPeriodHours
		*out = new(float64)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositoryObservation.
func (in *ChefRepositoryObservation) DeepCopy() *ChefRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(ChefRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositoryParameters) DeepCopyInto(out *ChefRepositoryParameters) {
	*out = *in
	if in.AllowAnyHostAuth != nil {
		in, out := &in.AllowAnyHostAuth, &out.AllowAnyHostAuth
		*out = new(bool)
		**out = **in
	}
	if in.AssumedOfflinePeriodSecs != nil {
		in, out := &in.AssumedOfflinePeriodSecs, &out.AssumedOfflinePeriodSecs
		*out = new(float64)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.BlockMismatchingMimeTypes != nil {
		in, out := &in.BlockMismatchingMimeTypes, &out.BlockMismatchingMimeTypes
		*out = new(bool)
		**out = **in
	}
	if in.BypassHeadRequests != nil {
		in, out := &in.BypassHeadRequests, &out.BypassHeadRequests
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.ClientTLSCertificate != nil {
		in, out := &in.ClientTLSCertificate, &out.ClientTLSCertificate
		*out = new(string)
		**out = **in
	}
	if in.ContentSynchronisation != nil {
		in, out := &in.ContentSynchronisation, &out.ContentSynchronisation
		*out = make([]ContentSynchronisationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableProxy != nil {
		in, out := &in.DisableProxy, &out.DisableProxy
		*out = new(bool)
		**out = **in
	}
	if in.DisableURLNormalization != nil {
		in, out := &in.DisableURLNormalization, &out.DisableURLNormalization
		*out = new(bool)
		**out = **in
	}
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.EnableCookieManagement != nil {
		in, out := &in.EnableCookieManagement, &out.EnableCookieManagement
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
		*out = new(string)
		**out = **in
	}
	if in.HardFail != nil {
		in, out := &in.HardFail, &out.HardFail
		*out = new(bool)
		**out = **in
	}
	if in.IncludesPattern != nil {
		in, out := &in.IncludesPattern, &out.IncludesPattern
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ListRemoteFolderItems != nil {
		in, out := &in.ListRemoteFolderItems, &out.ListRemoteFolderItems
		*out = new(bool)
		**out = **in
	}
	if in.LocalAddress != nil {
		in, out := &in.LocalAddress, &out.LocalAddress
		*out = new(string)
		**out = **in
	}
	if in.MetadataRetrievalTimeoutSecs != nil {
		in, out := &in.MetadataRetrievalTimeoutSecs, &out.MetadataRetrievalTimeoutSecs
		*out = new(float64)
		**out = **in
	}
	if in.MismatchingMimeTypesOverrideList != nil {
		in, out := &in.MismatchingMimeTypesOverrideList, &out.MismatchingMimeTypesOverrideList
		*out = new(string)
		**out = **in
	}
	if in.MissedCachePeriodSeconds != nil {
		in, out := &in.MissedCachePeriodSeconds, &out.MissedCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.Offline != nil {
		in, out := &in.Offline, &out.Offline
		*out = new(bool)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PriorityResolution != nil {
		in, out := &in.PriorityResolution, &out.PriorityResolution
		*out = new(bool)
		**out = **in
	}
	if in.ProjectEnvironments != nil {
		in, out := &in.ProjectEnvironments, &out.ProjectEnvironments
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.PropertySets != nil {
		in, out := &in.PropertySets, &out.PropertySets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = new(string)
		**out = **in
	}
	if in.RemoteRepoLayoutRef != nil {
		in, out := &in.RemoteRepoLayoutRef, &out.RemoteRepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.RetrievalCachePeriodSeconds != nil {
		in, out := &in.RetrievalCachePeriodSeconds, &out.RetrievalCachePeriodSeconds
		*out = new(float64)
		**out = **in
	}
	if in.ShareConfiguration != nil {
		in, out := &in.ShareConfiguration, &out.ShareConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.SocketTimeoutMillis != nil {
		in, out := &in.SocketTimeoutMillis, &out.SocketTimeoutMillis
		*out = new(float64)
		**out = **in
	}
	if in.StoreArtifactsLocally != nil {
		in, out := &in.StoreArtifactsLocally, &out.StoreArtifactsLocally
		*out = new(bool)
		**out = **in
	}
	if in.SynchronizeProperties != nil {
		in, out := &in.SynchronizeProperties, &out.SynchronizeProperties
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UnusedArtifactsCleanupPeriodHours != nil {
		in, out := &in.UnusedArtifactsCleanupPeriodHours, &out.UnusedArtifactsCleanupPeriodHours
		*out = new(float64)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositoryParameters.
func (in *ChefRepositoryParameters) DeepCopy() *ChefRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(ChefRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositorySpec) DeepCopyInto(out *ChefRepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositorySpec.
func (in *ChefRepositorySpec) DeepCopy() *ChefRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(ChefRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChefRepositoryStatus) DeepCopyInto(out *ChefRepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChefRepositoryStatus.
func (in *ChefRepositoryStatus) DeepCopy() *ChefRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(ChefRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentSynchronisationInitParameters) DeepCopyInto(out *ContentSynchronisationInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertiesEnabled != nil {
		in, out := &in.PropertiesEnabled, &out.PropertiesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SourceOriginAbsenceDetection != nil {
		in, out := &in.SourceOriginAbsenceDetection, &out.SourceOriginAbsenceDetection
		*out = new(bool)
		**out = **in
	}
	if in.StatisticsEnabled != nil {
		in, out := &in.StatisticsEnabled, &out.StatisticsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentSynchronisationInitParameters.
func (in *ContentSynchronisationInitParameters) DeepCopy() *ContentSynchronisationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ContentSynchronisationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentSynchronisationObservation) DeepCopyInto(out *ContentSynchronisationObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertiesEnabled != nil {
		in, out := &in.PropertiesEnabled, &out.PropertiesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SourceOriginAbsenceDetection != nil {
		in, out := &in.SourceOriginAbsenceDetection, &out.SourceOriginAbsenceDetection
		*out = new(bool)
		**out = **in
	}
	if in.StatisticsEnabled != nil {
		in, out := &in.StatisticsEnabled, &out.StatisticsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentSynchronisationObservation.
func (in *ContentSynchronisationObservation) DeepCopy() *ContentSynchronisationObservation {
	if in == nil {
		return nil
	}
	out := new(ContentSynchronisationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentSynchronisationParameters) DeepCopyInto(out *ContentSynchronisationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PropertiesEnabled != nil {
		in, out := &in.PropertiesEnabled, &out.PropertiesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SourceOriginAbsenceDetection != nil {
		in, out := &in.SourceOriginAbsenceDetection, &out.SourceOriginAbsenceDetection
		*out = new(bool)
		**out = **in
	}
	if in.StatisticsEnabled != nil {
		in, out := &in.StatisticsEnabled, &out.StatisticsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentSynchronisationParameters.
func (in *ContentSynchronisationParameters) DeepCopy() *ContentSynchronisationParameters {
	if in == nil {
		return nil
	}
	out := new(ContentSynchronisationParameters)
	in.DeepCopyInto(out)
	return out
}
