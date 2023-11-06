package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersEnvValueFromSecretKeyRef struct {
	// A Cloud Secret Manager secret version. Must be 'latest' for the latest version or an integer for a specific version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloud_run_service#key GoogleCloudRunService#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the secret in Cloud Secret Manager.
	//
	// By default, the secret is assumed to be in the same project.
	// If the secret is in another project, you must define an alias.
	// An alias definition has the form: :projects/{project-id|project-number}/secrets/.
	// If multiple alias definitions are needed, they must be separated by commas.
	// The alias definitions must be set on the run.googleapis.com/secrets annotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

