package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupAcceleratorConfigs struct {
	// Optional. Count of cores of this accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#core_count GoogleWorkbenchInstance#core_count}
	CoreCount *string `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Optional. Type of this accelerator. Possible values: ["NVIDIA_TESLA_P100", "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_A100", "NVIDIA_A100_80GB", "NVIDIA_L4", "NVIDIA_TESLA_T4_VWS", "NVIDIA_TESLA_P100_VWS", "NVIDIA_TESLA_P4_VWS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_workbench_instance#type GoogleWorkbenchInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

