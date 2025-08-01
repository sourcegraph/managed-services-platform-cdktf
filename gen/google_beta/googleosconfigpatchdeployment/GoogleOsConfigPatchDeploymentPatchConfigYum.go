package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigYum struct {
	// List of packages to exclude from update. These packages will be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#excludes GoogleOsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of packages to be updated.
	//
	// These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#exclusive_packages GoogleOsConfigPatchDeployment#exclusive_packages}
	ExclusivePackages *[]*string `field:"optional" json:"exclusivePackages" yaml:"exclusivePackages"`
	// Will cause patch to run yum update-minimal instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#minimal GoogleOsConfigPatchDeployment#minimal}
	Minimal interface{} `field:"optional" json:"minimal" yaml:"minimal"`
	// Adds the --security flag to yum update. Not supported on all platforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_patch_deployment#security GoogleOsConfigPatchDeployment#security}
	Security interface{} `field:"optional" json:"security" yaml:"security"`
}

