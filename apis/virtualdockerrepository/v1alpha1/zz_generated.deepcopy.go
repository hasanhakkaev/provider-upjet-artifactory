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
func (in *DockerRepository) DeepCopyInto(out *DockerRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepository.
func (in *DockerRepository) DeepCopy() *DockerRepository {
	if in == nil {
		return nil
	}
	out := new(DockerRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositoryInitParameters) DeepCopyInto(out *DockerRepositoryInitParameters) {
	*out = *in
	if in.ArtifactoryRequestsCanRetrieveRemoteArtifacts != nil {
		in, out := &in.ArtifactoryRequestsCanRetrieveRemoteArtifacts, &out.ArtifactoryRequestsCanRetrieveRemoteArtifacts
		*out = new(bool)
		**out = **in
	}
	if in.DefaultDeploymentRepo != nil {
		in, out := &in.DefaultDeploymentRepo, &out.DefaultDeploymentRepo
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResolveDockerTagsByTimestamp != nil {
		in, out := &in.ResolveDockerTagsByTimestamp, &out.ResolveDockerTagsByTimestamp
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositoryInitParameters.
func (in *DockerRepositoryInitParameters) DeepCopy() *DockerRepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(DockerRepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositoryList) DeepCopyInto(out *DockerRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositoryList.
func (in *DockerRepositoryList) DeepCopy() *DockerRepositoryList {
	if in == nil {
		return nil
	}
	out := new(DockerRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositoryObservation) DeepCopyInto(out *DockerRepositoryObservation) {
	*out = *in
	if in.ArtifactoryRequestsCanRetrieveRemoteArtifacts != nil {
		in, out := &in.ArtifactoryRequestsCanRetrieveRemoteArtifacts, &out.ArtifactoryRequestsCanRetrieveRemoteArtifacts
		*out = new(bool)
		**out = **in
	}
	if in.DefaultDeploymentRepo != nil {
		in, out := &in.DefaultDeploymentRepo, &out.DefaultDeploymentRepo
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResolveDockerTagsByTimestamp != nil {
		in, out := &in.ResolveDockerTagsByTimestamp, &out.ResolveDockerTagsByTimestamp
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositoryObservation.
func (in *DockerRepositoryObservation) DeepCopy() *DockerRepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(DockerRepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositoryParameters) DeepCopyInto(out *DockerRepositoryParameters) {
	*out = *in
	if in.ArtifactoryRequestsCanRetrieveRemoteArtifacts != nil {
		in, out := &in.ArtifactoryRequestsCanRetrieveRemoteArtifacts, &out.ArtifactoryRequestsCanRetrieveRemoteArtifacts
		*out = new(bool)
		**out = **in
	}
	if in.DefaultDeploymentRepo != nil {
		in, out := &in.DefaultDeploymentRepo, &out.DefaultDeploymentRepo
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
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
	if in.RepoLayoutRef != nil {
		in, out := &in.RepoLayoutRef, &out.RepoLayoutRef
		*out = new(string)
		**out = **in
	}
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResolveDockerTagsByTimestamp != nil {
		in, out := &in.ResolveDockerTagsByTimestamp, &out.ResolveDockerTagsByTimestamp
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositoryParameters.
func (in *DockerRepositoryParameters) DeepCopy() *DockerRepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(DockerRepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositorySpec) DeepCopyInto(out *DockerRepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositorySpec.
func (in *DockerRepositorySpec) DeepCopy() *DockerRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(DockerRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerRepositoryStatus) DeepCopyInto(out *DockerRepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerRepositoryStatus.
func (in *DockerRepositoryStatus) DeepCopy() *DockerRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(DockerRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
