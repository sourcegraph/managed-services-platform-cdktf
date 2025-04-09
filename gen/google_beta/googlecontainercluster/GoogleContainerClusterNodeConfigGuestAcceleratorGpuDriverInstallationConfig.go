package googlecontainercluster


type GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig struct {
	// Mode for how the GPU driver is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_container_cluster#gpu_driver_version GoogleContainerCluster#gpu_driver_version}
	GpuDriverVersion *string `field:"required" json:"gpuDriverVersion" yaml:"gpuDriverVersion"`
}

