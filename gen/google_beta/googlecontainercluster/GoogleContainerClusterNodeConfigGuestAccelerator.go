package googlecontainercluster


type GoogleContainerClusterNodeConfigGuestAccelerator struct {
	// The number of the accelerator cards exposed to an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#count GoogleContainerCluster#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#type GoogleContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// gpu_driver_installation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#gpu_driver_installation_config GoogleContainerCluster#gpu_driver_installation_config}
	GpuDriverInstallationConfig *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig `field:"optional" json:"gpuDriverInstallationConfig" yaml:"gpuDriverInstallationConfig"`
	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#gpu_partition_size GoogleContainerCluster#gpu_partition_size}
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// gpu_sharing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#gpu_sharing_config GoogleContainerCluster#gpu_sharing_config}
	GpuSharingConfig *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
}

