package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildSecret struct {
	// Cloud KMS key name to use to decrypt these envs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#kms_key_name GoogleCloudbuildTrigger#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Map of environment variable name to its encrypted value.
	//
	// Secret environment variables must be unique across all of a build's secrets,
	// and must be used by at least one build step. Values can be at most 64 KB in size.
	// There can be at most 100 secret values across all of a build's secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#secret_env GoogleCloudbuildTrigger#secret_env}
	SecretEnv *map[string]*string `field:"optional" json:"secretEnv" yaml:"secretEnv"`
}

