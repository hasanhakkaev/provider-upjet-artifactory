//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1Repository) DeepCopyInto(out *DockerV1Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1Repository.
func (in *DockerV1Repository) DeepCopy() *DockerV1Repository {
	if in == nil {
		return nil
	}
	out := new(DockerV1Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerV1Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositoryInitParameters) DeepCopyInto(out *DockerV1RepositoryInitParameters) {
	*out = *in
	if in.ArchiveBrowsingEnabled != nil {
		in, out := &in.ArchiveBrowsingEnabled, &out.ArchiveBrowsingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.CleanupOnDelete != nil {
		in, out := &in.CleanupOnDelete, &out.CleanupOnDelete
		*out = new(bool)
		**out = **in
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
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
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
	if in.MaxUniqueTags != nil {
		in, out := &in.MaxUniqueTags, &out.MaxUniqueTags
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]MemberInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositoryInitParameters.
func (in *DockerV1RepositoryInitParameters) DeepCopy() *DockerV1RepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositoryList) DeepCopyInto(out *DockerV1RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerV1Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositoryList.
func (in *DockerV1RepositoryList) DeepCopy() *DockerV1RepositoryList {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerV1RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositoryObservation) DeepCopyInto(out *DockerV1RepositoryObservation) {
	*out = *in
	if in.APIVersion != nil {
		in, out := &in.APIVersion, &out.APIVersion
		*out = new(string)
		**out = **in
	}
	if in.ArchiveBrowsingEnabled != nil {
		in, out := &in.ArchiveBrowsingEnabled, &out.ArchiveBrowsingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.BlockPushingSchema1 != nil {
		in, out := &in.BlockPushingSchema1, &out.BlockPushingSchema1
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.CleanupOnDelete != nil {
		in, out := &in.CleanupOnDelete, &out.CleanupOnDelete
		*out = new(bool)
		**out = **in
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
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
		*out = new(string)
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
	if in.MaxUniqueTags != nil {
		in, out := &in.MaxUniqueTags, &out.MaxUniqueTags
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]MemberObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.TagRetention != nil {
		in, out := &in.TagRetention, &out.TagRetention
		*out = new(float64)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositoryObservation.
func (in *DockerV1RepositoryObservation) DeepCopy() *DockerV1RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositoryParameters) DeepCopyInto(out *DockerV1RepositoryParameters) {
	*out = *in
	if in.ArchiveBrowsingEnabled != nil {
		in, out := &in.ArchiveBrowsingEnabled, &out.ArchiveBrowsingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BlackedOut != nil {
		in, out := &in.BlackedOut, &out.BlackedOut
		*out = new(bool)
		**out = **in
	}
	if in.CdnRedirect != nil {
		in, out := &in.CdnRedirect, &out.CdnRedirect
		*out = new(bool)
		**out = **in
	}
	if in.CleanupOnDelete != nil {
		in, out := &in.CleanupOnDelete, &out.CleanupOnDelete
		*out = new(bool)
		**out = **in
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
	if in.DownloadDirect != nil {
		in, out := &in.DownloadDirect, &out.DownloadDirect
		*out = new(bool)
		**out = **in
	}
	if in.ExcludesPattern != nil {
		in, out := &in.ExcludesPattern, &out.ExcludesPattern
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
	if in.MaxUniqueTags != nil {
		in, out := &in.MaxUniqueTags, &out.MaxUniqueTags
		*out = new(float64)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]MemberParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.XrayIndex != nil {
		in, out := &in.XrayIndex, &out.XrayIndex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositoryParameters.
func (in *DockerV1RepositoryParameters) DeepCopy() *DockerV1RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositorySpec) DeepCopyInto(out *DockerV1RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositorySpec.
func (in *DockerV1RepositorySpec) DeepCopy() *DockerV1RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerV1RepositoryStatus) DeepCopyInto(out *DockerV1RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerV1RepositoryStatus.
func (in *DockerV1RepositoryStatus) DeepCopy() *DockerV1RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(DockerV1RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberInitParameters) DeepCopyInto(out *MemberInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberInitParameters.
func (in *MemberInitParameters) DeepCopy() *MemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberObservation) DeepCopyInto(out *MemberObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberObservation.
func (in *MemberObservation) DeepCopy() *MemberObservation {
	if in == nil {
		return nil
	}
	out := new(MemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberParameters) DeepCopyInto(out *MemberParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberParameters.
func (in *MemberParameters) DeepCopy() *MemberParameters {
	if in == nil {
		return nil
	}
	out := new(MemberParameters)
	in.DeepCopyInto(out)
	return out
}
