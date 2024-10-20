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
func (in *CriteriaInitParameters) DeepCopyInto(out *CriteriaInitParameters) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaInitParameters.
func (in *CriteriaInitParameters) DeepCopy() *CriteriaInitParameters {
	if in == nil {
		return nil
	}
	out := new(CriteriaInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaObservation) DeepCopyInto(out *CriteriaObservation) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaObservation.
func (in *CriteriaObservation) DeepCopy() *CriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(CriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaParameters) DeepCopyInto(out *CriteriaParameters) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaParameters.
func (in *CriteriaParameters) DeepCopy() *CriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(CriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandlerInitParameters) DeepCopyInto(out *HandlerInitParameters) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = new(string)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandlerInitParameters.
func (in *HandlerInitParameters) DeepCopy() *HandlerInitParameters {
	if in == nil {
		return nil
	}
	out := new(HandlerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandlerObservation) DeepCopyInto(out *HandlerObservation) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = new(string)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandlerObservation.
func (in *HandlerObservation) DeepCopy() *HandlerObservation {
	if in == nil {
		return nil
	}
	out := new(HandlerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandlerParameters) DeepCopyInto(out *HandlerParameters) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = new(string)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandlerParameters.
func (in *HandlerParameters) DeepCopy() *HandlerParameters {
	if in == nil {
		return nil
	}
	out := new(HandlerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhook) DeepCopyInto(out *ReleaseBundleCustomWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhook.
func (in *ReleaseBundleCustomWebhook) DeepCopy() *ReleaseBundleCustomWebhook {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseBundleCustomWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookInitParameters) DeepCopyInto(out *ReleaseBundleCustomWebhookInitParameters) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]CriteriaInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]HandlerInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookInitParameters.
func (in *ReleaseBundleCustomWebhookInitParameters) DeepCopy() *ReleaseBundleCustomWebhookInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookList) DeepCopyInto(out *ReleaseBundleCustomWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleaseBundleCustomWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookList.
func (in *ReleaseBundleCustomWebhookList) DeepCopy() *ReleaseBundleCustomWebhookList {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseBundleCustomWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookObservation) DeepCopyInto(out *ReleaseBundleCustomWebhookObservation) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]CriteriaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]HandlerObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookObservation.
func (in *ReleaseBundleCustomWebhookObservation) DeepCopy() *ReleaseBundleCustomWebhookObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookParameters) DeepCopyInto(out *ReleaseBundleCustomWebhookParameters) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]CriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]HandlerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookParameters.
func (in *ReleaseBundleCustomWebhookParameters) DeepCopy() *ReleaseBundleCustomWebhookParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookSpec) DeepCopyInto(out *ReleaseBundleCustomWebhookSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookSpec.
func (in *ReleaseBundleCustomWebhookSpec) DeepCopy() *ReleaseBundleCustomWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleCustomWebhookStatus) DeepCopyInto(out *ReleaseBundleCustomWebhookStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleCustomWebhookStatus.
func (in *ReleaseBundleCustomWebhookStatus) DeepCopy() *ReleaseBundleCustomWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleCustomWebhookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhook) DeepCopyInto(out *ReleaseBundleWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhook.
func (in *ReleaseBundleWebhook) DeepCopy() *ReleaseBundleWebhook {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseBundleWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookCriteriaInitParameters) DeepCopyInto(out *ReleaseBundleWebhookCriteriaInitParameters) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookCriteriaInitParameters.
func (in *ReleaseBundleWebhookCriteriaInitParameters) DeepCopy() *ReleaseBundleWebhookCriteriaInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookCriteriaInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookCriteriaObservation) DeepCopyInto(out *ReleaseBundleWebhookCriteriaObservation) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookCriteriaObservation.
func (in *ReleaseBundleWebhookCriteriaObservation) DeepCopy() *ReleaseBundleWebhookCriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookCriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookCriteriaParameters) DeepCopyInto(out *ReleaseBundleWebhookCriteriaParameters) {
	*out = *in
	if in.AnyReleaseBundle != nil {
		in, out := &in.AnyReleaseBundle, &out.AnyReleaseBundle
		*out = new(bool)
		**out = **in
	}
	if in.ExcludePatterns != nil {
		in, out := &in.ExcludePatterns, &out.ExcludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludePatterns != nil {
		in, out := &in.IncludePatterns, &out.IncludePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegisteredReleaseBundleNames != nil {
		in, out := &in.RegisteredReleaseBundleNames, &out.RegisteredReleaseBundleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookCriteriaParameters.
func (in *ReleaseBundleWebhookCriteriaParameters) DeepCopy() *ReleaseBundleWebhookCriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookCriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookHandlerInitParameters) DeepCopyInto(out *ReleaseBundleWebhookHandlerInitParameters) {
	*out = *in
	if in.CustomHTTPHeaders != nil {
		in, out := &in.CustomHTTPHeaders, &out.CustomHTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookHandlerInitParameters.
func (in *ReleaseBundleWebhookHandlerInitParameters) DeepCopy() *ReleaseBundleWebhookHandlerInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookHandlerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookHandlerObservation) DeepCopyInto(out *ReleaseBundleWebhookHandlerObservation) {
	*out = *in
	if in.CustomHTTPHeaders != nil {
		in, out := &in.CustomHTTPHeaders, &out.CustomHTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookHandlerObservation.
func (in *ReleaseBundleWebhookHandlerObservation) DeepCopy() *ReleaseBundleWebhookHandlerObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookHandlerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookHandlerParameters) DeepCopyInto(out *ReleaseBundleWebhookHandlerParameters) {
	*out = *in
	if in.CustomHTTPHeaders != nil {
		in, out := &in.CustomHTTPHeaders, &out.CustomHTTPHeaders
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookHandlerParameters.
func (in *ReleaseBundleWebhookHandlerParameters) DeepCopy() *ReleaseBundleWebhookHandlerParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookHandlerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookInitParameters) DeepCopyInto(out *ReleaseBundleWebhookInitParameters) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]ReleaseBundleWebhookCriteriaInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]ReleaseBundleWebhookHandlerInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookInitParameters.
func (in *ReleaseBundleWebhookInitParameters) DeepCopy() *ReleaseBundleWebhookInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookList) DeepCopyInto(out *ReleaseBundleWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleaseBundleWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookList.
func (in *ReleaseBundleWebhookList) DeepCopy() *ReleaseBundleWebhookList {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseBundleWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookObservation) DeepCopyInto(out *ReleaseBundleWebhookObservation) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]ReleaseBundleWebhookCriteriaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]ReleaseBundleWebhookHandlerObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookObservation.
func (in *ReleaseBundleWebhookObservation) DeepCopy() *ReleaseBundleWebhookObservation {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookParameters) DeepCopyInto(out *ReleaseBundleWebhookParameters) {
	*out = *in
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]ReleaseBundleWebhookCriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = make([]ReleaseBundleWebhookHandlerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookParameters.
func (in *ReleaseBundleWebhookParameters) DeepCopy() *ReleaseBundleWebhookParameters {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookSpec) DeepCopyInto(out *ReleaseBundleWebhookSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookSpec.
func (in *ReleaseBundleWebhookSpec) DeepCopy() *ReleaseBundleWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBundleWebhookStatus) DeepCopyInto(out *ReleaseBundleWebhookStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBundleWebhookStatus.
func (in *ReleaseBundleWebhookStatus) DeepCopy() *ReleaseBundleWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseBundleWebhookStatus)
	in.DeepCopyInto(out)
	return out
}
