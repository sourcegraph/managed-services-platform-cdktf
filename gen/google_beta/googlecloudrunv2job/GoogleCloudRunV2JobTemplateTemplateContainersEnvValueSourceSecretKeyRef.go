package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersEnvValueSourceSecretKeyRef struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloud_run_v2_job#secret GoogleCloudRunV2Job#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloud_run_v2_job#version GoogleCloudRunV2Job#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

