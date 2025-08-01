package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo struct {
	// Cloud Build 2nd gen repository in the format of 'projects/<project>/locations/<location>/connections/<connection>/repositories/<repository>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_custom_target_type#repository GoogleClouddeployCustomTargetType#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Relative path from the repository root to the Skaffold file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_custom_target_type#path GoogleClouddeployCustomTargetType#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Branch or tag to use when cloning the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_custom_target_type#ref GoogleClouddeployCustomTargetType#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

