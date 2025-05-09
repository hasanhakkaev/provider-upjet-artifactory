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
func (in *HuggingfacemlRepository) DeepCopyInto(out *HuggingfacemlRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepository.
func (in *HuggingfacemlRepository) DeepCopy() *HuggingfacemlRepository {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HuggingfacemlRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositoryInitParameters) DeepCopyInto(out *HuggingfacemlRepositoryInitParameters) {
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositoryInitParameters.
func (in *HuggingfacemlRepositoryInitParameters) DeepCopy() *HuggingfacemlRepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositoryList) DeepCopyInto(out *HuggingfacemlRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HuggingfacemlRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositoryList.
func (in *HuggingfacemlRepositoryList) DeepCopy() *HuggingfacemlRepositoryList {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HuggingfacemlRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositoryObservation) DeepCopyInto(out *HuggingfacemlRepositoryObservation) {
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositoryObservation.
func (in *HuggingfacemlRepositoryObservation) DeepCopy() *HuggingfacemlRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositoryParameters) DeepCopyInto(out *HuggingfacemlRepositoryParameters) {
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositoryParameters.
func (in *HuggingfacemlRepositoryParameters) DeepCopy() *HuggingfacemlRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositorySpec) DeepCopyInto(out *HuggingfacemlRepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositorySpec.
func (in *HuggingfacemlRepositorySpec) DeepCopy() *HuggingfacemlRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HuggingfacemlRepositoryStatus) DeepCopyInto(out *HuggingfacemlRepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HuggingfacemlRepositoryStatus.
func (in *HuggingfacemlRepositoryStatus) DeepCopy() *HuggingfacemlRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(HuggingfacemlRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
