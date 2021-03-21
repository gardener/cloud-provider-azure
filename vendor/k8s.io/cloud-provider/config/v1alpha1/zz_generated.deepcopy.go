// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
func (in *CloudControllerManagerConfiguration) DeepCopyInto(out *CloudControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Generic.DeepCopyInto(&out.Generic)
	in.KubeCloudShared.DeepCopyInto(&out.KubeCloudShared)
	out.ServiceController = in.ServiceController
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudControllerManagerConfiguration.
func (in *CloudControllerManagerConfiguration) DeepCopy() *CloudControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderConfiguration) DeepCopyInto(out *CloudProviderConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderConfiguration.
func (in *CloudProviderConfiguration) DeepCopy() *CloudProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCloudSharedConfiguration) DeepCopyInto(out *KubeCloudSharedConfiguration) {
	*out = *in
	out.CloudProvider = in.CloudProvider
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	if in.ConfigureCloudRoutes != nil {
		in, out := &in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes
		*out = new(bool)
		**out = **in
	}
	out.NodeSyncPeriod = in.NodeSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCloudSharedConfiguration.
func (in *KubeCloudSharedConfiguration) DeepCopy() *KubeCloudSharedConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeCloudSharedConfiguration)
	in.DeepCopyInto(out)
	return out
}