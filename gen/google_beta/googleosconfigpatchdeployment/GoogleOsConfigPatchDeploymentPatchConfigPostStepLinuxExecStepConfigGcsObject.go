package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject struct {
	// Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#bucket GoogleOsConfigPatchDeployment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Generation number of the Cloud Storage object.
	//
	// This is used to ensure that the ExecStep specified by this PatchJob does not change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#generation_number GoogleOsConfigPatchDeployment#generation_number}
	GenerationNumber *string `field:"required" json:"generationNumber" yaml:"generationNumber"`
	// Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#object GoogleOsConfigPatchDeployment#object}
	Object *string `field:"required" json:"object" yaml:"object"`
}

