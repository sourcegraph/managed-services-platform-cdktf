package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersEnvValueSourceSecretKeyRef struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secretName} if the secret is in the same project. projects/{project}/secrets/{secretName} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloud_run_v2_service#secret GoogleCloudRunV2Service#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest' for the latest value or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloud_run_v2_service#version GoogleCloudRunV2Service#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

