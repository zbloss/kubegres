// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubegres) DeepCopyInto(out *Kubegres) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubegres.
func (in *Kubegres) DeepCopy() *Kubegres {
	if in == nil {
		return nil
	}
	out := new(Kubegres)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kubegres) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresBackUp) DeepCopyInto(out *KubegresBackUp) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresBackUp.
func (in *KubegresBackUp) DeepCopy() *KubegresBackUp {
	if in == nil {
		return nil
	}
	out := new(KubegresBackUp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresBlockingOperation) DeepCopyInto(out *KubegresBlockingOperation) {
	*out = *in
	out.StatefulSetOperation = in.StatefulSetOperation
	out.StatefulSetSpecUpdateOperation = in.StatefulSetSpecUpdateOperation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresBlockingOperation.
func (in *KubegresBlockingOperation) DeepCopy() *KubegresBlockingOperation {
	if in == nil {
		return nil
	}
	out := new(KubegresBlockingOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresDatabase) DeepCopyInto(out *KubegresDatabase) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresDatabase.
func (in *KubegresDatabase) DeepCopy() *KubegresDatabase {
	if in == nil {
		return nil
	}
	out := new(KubegresDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresFailover) DeepCopyInto(out *KubegresFailover) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresFailover.
func (in *KubegresFailover) DeepCopy() *KubegresFailover {
	if in == nil {
		return nil
	}
	out := new(KubegresFailover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresList) DeepCopyInto(out *KubegresList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kubegres, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresList.
func (in *KubegresList) DeepCopy() *KubegresList {
	if in == nil {
		return nil
	}
	out := new(KubegresList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubegresList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresSpec) DeepCopyInto(out *KubegresSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.Database.DeepCopyInto(&out.Database)
	out.Failover = in.Failover
	out.Backup = in.Backup
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresSpec.
func (in *KubegresSpec) DeepCopy() *KubegresSpec {
	if in == nil {
		return nil
	}
	out := new(KubegresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresStatefulSetOperation) DeepCopyInto(out *KubegresStatefulSetOperation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresStatefulSetOperation.
func (in *KubegresStatefulSetOperation) DeepCopy() *KubegresStatefulSetOperation {
	if in == nil {
		return nil
	}
	out := new(KubegresStatefulSetOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresStatefulSetSpecUpdateOperation) DeepCopyInto(out *KubegresStatefulSetSpecUpdateOperation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresStatefulSetSpecUpdateOperation.
func (in *KubegresStatefulSetSpecUpdateOperation) DeepCopy() *KubegresStatefulSetSpecUpdateOperation {
	if in == nil {
		return nil
	}
	out := new(KubegresStatefulSetSpecUpdateOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubegresStatus) DeepCopyInto(out *KubegresStatus) {
	*out = *in
	out.BlockingOperation = in.BlockingOperation
	out.PreviousBlockingOperation = in.PreviousBlockingOperation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubegresStatus.
func (in *KubegresStatus) DeepCopy() *KubegresStatus {
	if in == nil {
		return nil
	}
	out := new(KubegresStatus)
	in.DeepCopyInto(out)
	return out
}
