package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate struct {
	// Only apply updates of these windows update classifications.
	//
	// If empty, all updates are applied. Possible values: ["CRITICAL", "SECURITY", "DEFINITION", "DRIVER", "FEATURE_PACK", "SERVICE_PACK", "TOOL", "UPDATE_ROLLUP", "UPDATE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#classifications GoogleOsConfigPatchDeployment#classifications}
	Classifications *[]*string `field:"optional" json:"classifications" yaml:"classifications"`
	// List of KBs to exclude from update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#excludes GoogleOsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of kbs to be updated.
	//
	// These are the only patches that will be updated.
	// This field must not be used with other patch configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#exclusive_patches GoogleOsConfigPatchDeployment#exclusive_patches}
	ExclusivePatches *[]*string `field:"optional" json:"exclusivePatches" yaml:"exclusivePatches"`
}

