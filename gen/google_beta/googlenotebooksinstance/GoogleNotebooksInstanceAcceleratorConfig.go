package googlenotebooksinstance


type GoogleNotebooksInstanceAcceleratorConfig struct {
	// Count of cores of this accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_instance#core_count GoogleNotebooksInstance#core_count}
	CoreCount *float64 `field:"required" json:"coreCount" yaml:"coreCount"`
	// Type of this accelerator. Possible values: ["ACCELERATOR_TYPE_UNSPECIFIED", "NVIDIA_TESLA_K80", "NVIDIA_TESLA_P100", "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_T4_VWS", "NVIDIA_TESLA_P100_VWS", "NVIDIA_TESLA_P4_VWS", "NVIDIA_TESLA_A100", "TPU_V2", "TPU_V3"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_instance#type GoogleNotebooksInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

