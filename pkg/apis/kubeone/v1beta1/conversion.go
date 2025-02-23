/*
Copyright 2020 The KubeOne Authors.

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

package v1beta1

import (
	unsafe "unsafe"

	kubeoneapi "k8c.io/kubeone/pkg/apis/kubeone"

	conversion "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/utils/pointer"
)

func Convert_kubeone_ContainerRuntimeContainerd_To_v1beta1_ContainerRuntimeContainerd(*kubeoneapi.ContainerRuntimeContainerd, *ContainerRuntimeContainerd, conversion.Scope) error {
	// Skip conversion
	return nil
}

func Convert_kubeone_ContainerRuntimeDocker_To_v1beta1_ContainerRuntimeDocker(*kubeoneapi.ContainerRuntimeDocker, *ContainerRuntimeDocker, conversion.Scope) error {
	// Skip conversion
	return nil
}

// Convert_v1beta1_Features_To_kubeone_Features is an autogenerated conversion function.
func Convert_v1beta1_Features_To_kubeone_Features(in *Features, out *kubeoneapi.Features, s conversion.Scope) error {
	if err := autoConvert_v1beta1_Features_To_kubeone_Features(in, out, s); err != nil {
		return err
	}

	out.CoreDNS = &kubeoneapi.CoreDNS{
		Replicas:                  pointer.Int32(2),
		DeployPodDisruptionBudget: pointer.Bool(true),
	}

	// The PodPresets field has been dropped from v1beta2 API.
	return nil
}

// Convert_v1beta1_CloudProviderSpec_To_kubeone_CloudProviderSpec is an autogenerated conversion function.
func Convert_v1beta1_CloudProviderSpec_To_kubeone_CloudProviderSpec(in *CloudProviderSpec, out *kubeoneapi.CloudProviderSpec, s conversion.Scope) error {
	if err := autoConvert_v1beta1_CloudProviderSpec_To_kubeone_CloudProviderSpec(in, out, s); err != nil {
		return err
	}

	// PacketSpec has been renamed to EquinixMetalSpec
	out.EquinixMetal = (*kubeoneapi.EquinixMetalSpec)(unsafe.Pointer(in.Packet))

	return nil
}

// Convert_kubeone_CloudProviderSpec_To_v1beta1_CloudProviderSpec is an autogenerated conversion function.
func Convert_kubeone_CloudProviderSpec_To_v1beta1_CloudProviderSpec(in *kubeoneapi.CloudProviderSpec, out *CloudProviderSpec, s conversion.Scope) error {
	if err := autoConvert_kubeone_CloudProviderSpec_To_v1beta1_CloudProviderSpec(in, out, s); err != nil {
		return err
	}

	// PacketSpec has been renamed to EquinixMetalSpec
	out.Packet = (*PacketSpec)(unsafe.Pointer(in.EquinixMetal))

	return nil
}

func Convert_kubeone_HostConfig_To_v1beta1_HostConfig(in *kubeoneapi.HostConfig, out *HostConfig, scope conversion.Scope) error {
	// explicitly skip kubelet conversion omitted in autoConvert_kubeone_HostConfig_To_v1beta1_HostConfig
	return autoConvert_kubeone_HostConfig_To_v1beta1_HostConfig(in, out, scope)
}

func Convert_kubeone_KubeOneCluster_To_v1beta1_KubeOneCluster(in *kubeoneapi.KubeOneCluster, out *KubeOneCluster, s conversion.Scope) error {
	// LoggingConfig was introduced only in new v1beta2 API, so we skip it here
	return autoConvert_kubeone_KubeOneCluster_To_v1beta1_KubeOneCluster(in, out, s)
}

func Convert_kubeone_ProviderSpec_To_v1beta1_ProviderSpec(in *kubeoneapi.ProviderSpec, out *ProviderSpec, s conversion.Scope) error {
	// NodeAnnotations and MachineObjectAnnotations were introduced only in new v1beta2 API, so we skip them here
	return autoConvert_kubeone_ProviderSpec_To_v1beta1_ProviderSpec(in, out, s)
}

func Convert_kubeone_Features_To_v1beta1_Features(in *kubeoneapi.Features, out *Features, s conversion.Scope) error {
	// CoreDNS feature is introduced only in the v1beta2 API
	return autoConvert_kubeone_Features_To_v1beta1_Features(in, out, s)
}
