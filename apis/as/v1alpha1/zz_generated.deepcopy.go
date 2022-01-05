//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 NDD.

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
	"github.com/yndd/nddo-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AsRegister) DeepCopyInto(out *AsRegister) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make([]*v1.Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SourceTag != nil {
		in, out := &in.SourceTag, &out.SourceTag
		*out = make([]*v1.Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AsRegister.
func (in *AsRegister) DeepCopy() *AsRegister {
	if in == nil {
		return nil
	}
	out := new(AsRegister)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NddrAsRegister) DeepCopyInto(out *NddrAsRegister) {
	*out = *in
	in.AsRegister.DeepCopyInto(&out.AsRegister)
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(NddrRegisterState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NddrAsRegister.
func (in *NddrAsRegister) DeepCopy() *NddrAsRegister {
	if in == nil {
		return nil
	}
	out := new(NddrAsRegister)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NddrRegisterState) DeepCopyInto(out *NddrRegisterState) {
	*out = *in
	if in.As != nil {
		in, out := &in.As, &out.As
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NddrRegisterState.
func (in *NddrRegisterState) DeepCopy() *NddrRegisterState {
	if in == nil {
		return nil
	}
	out := new(NddrRegisterState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NddrRegistry) DeepCopyInto(out *NddrRegistry) {
	*out = *in
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = make([]*NddrRegistryRegistry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NddrRegistryRegistry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NddrRegistry.
func (in *NddrRegistry) DeepCopy() *NddrRegistry {
	if in == nil {
		return nil
	}
	out := new(NddrRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NddrRegistryRegistry) DeepCopyInto(out *NddrRegistryRegistry) {
	*out = *in
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.AllocationStrategy != nil {
		in, out := &in.AllocationStrategy, &out.AllocationStrategy
		*out = new(string)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(uint32)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(uint32)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(NddrRegistryRegistryState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NddrRegistryRegistry.
func (in *NddrRegistryRegistry) DeepCopy() *NddrRegistryRegistry {
	if in == nil {
		return nil
	}
	out := new(NddrRegistryRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NddrRegistryRegistryState) DeepCopyInto(out *NddrRegistryRegistryState) {
	*out = *in
	if in.Total != nil {
		in, out := &in.Total, &out.Total
		*out = new(uint32)
		**out = **in
	}
	if in.Allocated != nil {
		in, out := &in.Allocated, &out.Allocated
		*out = new(uint32)
		**out = **in
	}
	if in.Available != nil {
		in, out := &in.Available, &out.Available
		*out = new(uint32)
		**out = **in
	}
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = make([]*uint32, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(uint32)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NddrRegistryRegistryState.
func (in *NddrRegistryRegistryState) DeepCopy() *NddrRegistryRegistryState {
	if in == nil {
		return nil
	}
	out := new(NddrRegistryRegistryState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Register) DeepCopyInto(out *Register) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Register.
func (in *Register) DeepCopy() *Register {
	if in == nil {
		return nil
	}
	out := new(Register)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Register) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterList) DeepCopyInto(out *RegisterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Register, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterList.
func (in *RegisterList) DeepCopy() *RegisterList {
	if in == nil {
		return nil
	}
	out := new(RegisterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegisterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterSpec) DeepCopyInto(out *RegisterSpec) {
	*out = *in
	in.OdaInfo.DeepCopyInto(&out.OdaInfo)
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Register != nil {
		in, out := &in.Register, &out.Register
		*out = new(AsRegister)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterSpec.
func (in *RegisterSpec) DeepCopy() *RegisterSpec {
	if in == nil {
		return nil
	}
	out := new(RegisterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegisterStatus) DeepCopyInto(out *RegisterStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	in.OdaInfo.DeepCopyInto(&out.OdaInfo)
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Register != nil {
		in, out := &in.Register, &out.Register
		*out = new(NddrAsRegister)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegisterStatus.
func (in *RegisterStatus) DeepCopy() *RegisterStatus {
	if in == nil {
		return nil
	}
	out := new(RegisterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryRegistry) DeepCopyInto(out *RegistryRegistry) {
	*out = *in
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.AllocationStrategy != nil {
		in, out := &in.AllocationStrategy, &out.AllocationStrategy
		*out = new(string)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(uint32)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(uint32)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryRegistry.
func (in *RegistryRegistry) DeepCopy() *RegistryRegistry {
	if in == nil {
		return nil
	}
	out := new(RegistryRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.OdaInfo.DeepCopyInto(&out.OdaInfo)
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(RegistryRegistry)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	in.OdaInfo.DeepCopyInto(&out.OdaInfo)
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(NddrRegistryRegistry)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Root) DeepCopyInto(out *Root) {
	*out = *in
	if in.RegistryNddrRegistry != nil {
		in, out := &in.RegistryNddrRegistry, &out.RegistryNddrRegistry
		*out = new(NddrRegistry)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Root.
func (in *Root) DeepCopy() *Root {
	if in == nil {
		return nil
	}
	out := new(Root)
	in.DeepCopyInto(out)
	return out
}
