package googleclouddeploycustomtargettype


type GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModules struct {
	// The Skaffold Config modules to use from the specified source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_clouddeploy_custom_target_type#configs GoogleClouddeployCustomTargetType#configs}
	Configs *[]*string `field:"optional" json:"configs" yaml:"configs"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_clouddeploy_custom_target_type#git GoogleClouddeployCustomTargetType#git}
	Git *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit `field:"optional" json:"git" yaml:"git"`
	// google_cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_clouddeploy_custom_target_type#google_cloud_storage GoogleClouddeployCustomTargetType#google_cloud_storage}
	GoogleCloudStorage *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage `field:"optional" json:"googleCloudStorage" yaml:"googleCloudStorage"`
}
