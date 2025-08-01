package containercluster


type ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig struct {
	// Mode for how the GPU driver is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#gpu_driver_version ContainerCluster#gpu_driver_version}
	GpuDriverVersion *string `field:"required" json:"gpuDriverVersion" yaml:"gpuDriverVersion"`
}

