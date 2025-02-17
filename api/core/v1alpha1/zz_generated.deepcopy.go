//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavor) DeepCopyInto(out *Flavor) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavor.
func (in *Flavor) DeepCopy() *Flavor {
	if in == nil {
		return nil
	}
	out := new(Flavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceConfig) DeepCopyInto(out *InferenceConfig) {
	*out = *in
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]Flavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceConfig.
func (in *InferenceConfig) DeepCopy() *InferenceConfig {
	if in == nil {
		return nil
	}
	out := new(InferenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelClaim) DeepCopyInto(out *ModelClaim) {
	*out = *in
	if in.InferenceFlavors != nil {
		in, out := &in.InferenceFlavors, &out.InferenceFlavors
		*out = make([]FlavorName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelClaim.
func (in *ModelClaim) DeepCopy() *ModelClaim {
	if in == nil {
		return nil
	}
	out := new(ModelClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelClaims) DeepCopyInto(out *ModelClaims) {
	*out = *in
	if in.Models != nil {
		in, out := &in.Models, &out.Models
		*out = make([]ModelRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InferenceFlavors != nil {
		in, out := &in.InferenceFlavors, &out.InferenceFlavors
		*out = make([]FlavorName, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelClaims.
func (in *ModelClaims) DeepCopy() *ModelClaims {
	if in == nil {
		return nil
	}
	out := new(ModelClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelHub) DeepCopyInto(out *ModelHub) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	if in.AllowPatterns != nil {
		in, out := &in.AllowPatterns, &out.AllowPatterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnorePatterns != nil {
		in, out := &in.IgnorePatterns, &out.IgnorePatterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelHub.
func (in *ModelHub) DeepCopy() *ModelHub {
	if in == nil {
		return nil
	}
	out := new(ModelHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelRef) DeepCopyInto(out *ModelRef) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(ModelRole)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelRef.
func (in *ModelRef) DeepCopy() *ModelRef {
	if in == nil {
		return nil
	}
	out := new(ModelRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSource) DeepCopyInto(out *ModelSource) {
	*out = *in
	if in.ModelHub != nil {
		in, out := &in.ModelHub, &out.ModelHub
		*out = new(ModelHub)
		(*in).DeepCopyInto(*out)
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(URIProtocol)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSource.
func (in *ModelSource) DeepCopy() *ModelSource {
	if in == nil {
		return nil
	}
	out := new(ModelSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.InferenceConfig != nil {
		in, out := &in.InferenceConfig, &out.InferenceConfig
		*out = new(InferenceConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelStatus) DeepCopyInto(out *ModelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelStatus.
func (in *ModelStatus) DeepCopy() *ModelStatus {
	if in == nil {
		return nil
	}
	out := new(ModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenModel) DeepCopyInto(out *OpenModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenModel.
func (in *OpenModel) DeepCopy() *OpenModel {
	if in == nil {
		return nil
	}
	out := new(OpenModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenModelList) DeepCopyInto(out *OpenModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenModelList.
func (in *OpenModelList) DeepCopy() *OpenModelList {
	if in == nil {
		return nil
	}
	out := new(OpenModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
