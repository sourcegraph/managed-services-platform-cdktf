package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateVolumesSecret struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#secret GoogleCloudRunV2Service#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Integer representation of mode bits to use on created files by default.
	//
	// Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#default_mode GoogleCloudRunV2Service#default_mode}
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#items GoogleCloudRunV2Service#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

