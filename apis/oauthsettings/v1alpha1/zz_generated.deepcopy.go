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
func (in *OauthProviderInitParameters) DeepCopyInto(out *OauthProviderInitParameters) {
	*out = *in
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.AuthURL != nil {
		in, out := &in.AuthURL, &out.AuthURL
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	out.ClientSecretSecretRef = in.ClientSecretSecretRef
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthProviderInitParameters.
func (in *OauthProviderInitParameters) DeepCopy() *OauthProviderInitParameters {
	if in == nil {
		return nil
	}
	out := new(OauthProviderInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthProviderObservation) DeepCopyInto(out *OauthProviderObservation) {
	*out = *in
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.AuthURL != nil {
		in, out := &in.AuthURL, &out.AuthURL
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthProviderObservation.
func (in *OauthProviderObservation) DeepCopy() *OauthProviderObservation {
	if in == nil {
		return nil
	}
	out := new(OauthProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthProviderParameters) DeepCopyInto(out *OauthProviderParameters) {
	*out = *in
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.AuthURL != nil {
		in, out := &in.AuthURL, &out.AuthURL
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	out.ClientSecretSecretRef = in.ClientSecretSecretRef
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TokenURL != nil {
		in, out := &in.TokenURL, &out.TokenURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthProviderParameters.
func (in *OauthProviderParameters) DeepCopy() *OauthProviderParameters {
	if in == nil {
		return nil
	}
	out := new(OauthProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Settings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsInitParameters) DeepCopyInto(out *SettingsInitParameters) {
	*out = *in
	if in.AllowUserToAccessProfile != nil {
		in, out := &in.AllowUserToAccessProfile, &out.AllowUserToAccessProfile
		*out = new(bool)
		**out = **in
	}
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.OauthProvider != nil {
		in, out := &in.OauthProvider, &out.OauthProvider
		*out = make([]OauthProviderInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistUsers != nil {
		in, out := &in.PersistUsers, &out.PersistUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsInitParameters.
func (in *SettingsInitParameters) DeepCopy() *SettingsInitParameters {
	if in == nil {
		return nil
	}
	out := new(SettingsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsList) DeepCopyInto(out *SettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Settings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsList.
func (in *SettingsList) DeepCopy() *SettingsList {
	if in == nil {
		return nil
	}
	out := new(SettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsObservation) DeepCopyInto(out *SettingsObservation) {
	*out = *in
	if in.AllowUserToAccessProfile != nil {
		in, out := &in.AllowUserToAccessProfile, &out.AllowUserToAccessProfile
		*out = new(bool)
		**out = **in
	}
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.OauthProvider != nil {
		in, out := &in.OauthProvider, &out.OauthProvider
		*out = make([]OauthProviderObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistUsers != nil {
		in, out := &in.PersistUsers, &out.PersistUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsObservation.
func (in *SettingsObservation) DeepCopy() *SettingsObservation {
	if in == nil {
		return nil
	}
	out := new(SettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsParameters) DeepCopyInto(out *SettingsParameters) {
	*out = *in
	if in.AllowUserToAccessProfile != nil {
		in, out := &in.AllowUserToAccessProfile, &out.AllowUserToAccessProfile
		*out = new(bool)
		**out = **in
	}
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.OauthProvider != nil {
		in, out := &in.OauthProvider, &out.OauthProvider
		*out = make([]OauthProviderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistUsers != nil {
		in, out := &in.PersistUsers, &out.PersistUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsParameters.
func (in *SettingsParameters) DeepCopy() *SettingsParameters {
	if in == nil {
		return nil
	}
	out := new(SettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsSpec) DeepCopyInto(out *SettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsSpec.
func (in *SettingsSpec) DeepCopy() *SettingsSpec {
	if in == nil {
		return nil
	}
	out := new(SettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsStatus) DeepCopyInto(out *SettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsStatus.
func (in *SettingsStatus) DeepCopy() *SettingsStatus {
	if in == nil {
		return nil
	}
	out := new(SettingsStatus)
	in.DeepCopyInto(out)
	return out
}
