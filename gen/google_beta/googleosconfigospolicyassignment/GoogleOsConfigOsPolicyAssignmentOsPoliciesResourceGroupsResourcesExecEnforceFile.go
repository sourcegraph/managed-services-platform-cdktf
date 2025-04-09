package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFile struct {
	// Defaults to false.
	//
	// When false, files are subject to validations based on the file type: Remote: A checksum must be specified. Cloud Storage: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_os_policy_assignment#allow_insecure GoogleOsConfigOsPolicyAssignment#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_os_policy_assignment#gcs GoogleOsConfigOsPolicyAssignment#gcs}
	Gcs *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// A local path within the VM to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_os_policy_assignment#local_path GoogleOsConfigOsPolicyAssignment#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_os_policy_assignment#remote GoogleOsConfigOsPolicyAssignment#remote}
	Remote *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileRemote `field:"optional" json:"remote" yaml:"remote"`
}

