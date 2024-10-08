//go:build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaUsageDetails) DeepCopyInto(out *QuotaUsageDetails) {
	*out = *in
	if in.Reserved != nil {
		in, out := &in.Reserved, &out.Reserved
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaUsageDetails.
func (in *QuotaUsageDetails) DeepCopy() *QuotaUsageDetails {
	if in == nil {
		return nil
	}
	out := new(QuotaUsageDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTypeLevelQuotaStatus) DeepCopyInto(out *ResourceTypeLevelQuotaStatus) {
	*out = *in
	if in.ResourceTypeSCLevelQuotaStatuses != nil {
		in, out := &in.ResourceTypeSCLevelQuotaStatuses, &out.ResourceTypeSCLevelQuotaStatuses
		*out = make(SCLevelQuotaStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTypeLevelQuotaStatus.
func (in *ResourceTypeLevelQuotaStatus) DeepCopy() *ResourceTypeLevelQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceTypeLevelQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ResourceTypeLevelQuotaStatusList) DeepCopyInto(out *ResourceTypeLevelQuotaStatusList) {
	{
		in := &in
		*out = make(ResourceTypeLevelQuotaStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTypeLevelQuotaStatusList.
func (in ResourceTypeLevelQuotaStatusList) DeepCopy() ResourceTypeLevelQuotaStatusList {
	if in == nil {
		return nil
	}
	out := new(ResourceTypeLevelQuotaStatusList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCLevelQuotaStatus) DeepCopyInto(out *SCLevelQuotaStatus) {
	*out = *in
	if in.SCLevelQuotaUsage != nil {
		in, out := &in.SCLevelQuotaUsage, &out.SCLevelQuotaUsage
		*out = new(QuotaUsageDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCLevelQuotaStatus.
func (in *SCLevelQuotaStatus) DeepCopy() *SCLevelQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(SCLevelQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SCLevelQuotaStatusList) DeepCopyInto(out *SCLevelQuotaStatusList) {
	{
		in := &in
		*out = make(SCLevelQuotaStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCLevelQuotaStatusList.
func (in SCLevelQuotaStatusList) DeepCopy() SCLevelQuotaStatusList {
	if in == nil {
		return nil
	}
	out := new(SCLevelQuotaStatusList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyQuota) DeepCopyInto(out *StoragePolicyQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyQuota.
func (in *StoragePolicyQuota) DeepCopy() *StoragePolicyQuota {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicyQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyQuotaList) DeepCopyInto(out *StoragePolicyQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePolicyQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyQuotaList.
func (in *StoragePolicyQuotaList) DeepCopy() *StoragePolicyQuotaList {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicyQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyQuotaSpec) DeepCopyInto(out *StoragePolicyQuotaSpec) {
	*out = *in
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyQuotaSpec.
func (in *StoragePolicyQuotaSpec) DeepCopy() *StoragePolicyQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyQuotaStatus) DeepCopyInto(out *StoragePolicyQuotaStatus) {
	*out = *in
	if in.SCLevelQuotaStatuses != nil {
		in, out := &in.SCLevelQuotaStatuses, &out.SCLevelQuotaStatuses
		*out = make(SCLevelQuotaStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceTypeLevelQuotaStatuses != nil {
		in, out := &in.ResourceTypeLevelQuotaStatuses, &out.ResourceTypeLevelQuotaStatuses
		*out = make(ResourceTypeLevelQuotaStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyQuotaStatus.
func (in *StoragePolicyQuotaStatus) DeepCopy() *StoragePolicyQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyUsage) DeepCopyInto(out *StoragePolicyUsage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyUsage.
func (in *StoragePolicyUsage) DeepCopy() *StoragePolicyUsage {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicyUsage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyUsageList) DeepCopyInto(out *StoragePolicyUsageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePolicyUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyUsageList.
func (in *StoragePolicyUsageList) DeepCopy() *StoragePolicyUsageList {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyUsageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePolicyUsageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyUsageSpec) DeepCopyInto(out *StoragePolicyUsageSpec) {
	*out = *in
	if in.ResourceAPIgroup != nil {
		in, out := &in.ResourceAPIgroup, &out.ResourceAPIgroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyUsageSpec.
func (in *StoragePolicyUsageSpec) DeepCopy() *StoragePolicyUsageSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyUsageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePolicyUsageStatus) DeepCopyInto(out *StoragePolicyUsageStatus) {
	*out = *in
	if in.ResourceTypeLevelQuotaUsage != nil {
		in, out := &in.ResourceTypeLevelQuotaUsage, &out.ResourceTypeLevelQuotaUsage
		*out = new(QuotaUsageDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePolicyUsageStatus.
func (in *StoragePolicyUsageStatus) DeepCopy() *StoragePolicyUsageStatus {
	if in == nil {
		return nil
	}
	out := new(StoragePolicyUsageStatus)
	in.DeepCopyInto(out)
	return out
}
