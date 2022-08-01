//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contract) DeepCopyInto(out *Contract) {
	*out = *in
	in.ExpiryTimestamp.DeepCopyInto(&out.ExpiryTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contract.
func (in *Contract) DeepCopy() *Contract {
	if in == nil {
		return nil
	}
	out := new(Contract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseRequest) DeepCopyInto(out *LicenseRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(LicenseRequestRequest)
		**out = **in
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(LicenseRequestResponse)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseRequest.
func (in *LicenseRequest) DeepCopy() *LicenseRequest {
	if in == nil {
		return nil
	}
	out := new(LicenseRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseRequestRequest) DeepCopyInto(out *LicenseRequestRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseRequestRequest.
func (in *LicenseRequestRequest) DeepCopy() *LicenseRequestRequest {
	if in == nil {
		return nil
	}
	out := new(LicenseRequestRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseRequestResponse) DeepCopyInto(out *LicenseRequestResponse) {
	*out = *in
	in.Contract.DeepCopyInto(&out.Contract)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseRequestResponse.
func (in *LicenseRequestResponse) DeepCopy() *LicenseRequestResponse {
	if in == nil {
		return nil
	}
	out := new(LicenseRequestResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseStatus) DeepCopyInto(out *LicenseStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseStatus.
func (in *LicenseStatus) DeepCopy() *LicenseStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseStatusList) DeepCopyInto(out *LicenseStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseStatusList.
func (in *LicenseStatusList) DeepCopy() *LicenseStatusList {
	if in == nil {
		return nil
	}
	out := new(LicenseStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseStatusSpec) DeepCopyInto(out *LicenseStatusSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseStatusSpec.
func (in *LicenseStatusSpec) DeepCopy() *LicenseStatusSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseStatusStatus) DeepCopyInto(out *LicenseStatusStatus) {
	*out = *in
	in.License.DeepCopyInto(&out.License)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseStatusStatus.
func (in *LicenseStatusStatus) DeepCopy() *LicenseStatusStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseStatusStatus)
	in.DeepCopyInto(out)
	return out
}
