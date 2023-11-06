package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionServiceConfigSecretVolumes struct {
	// The path within the container to mount the secret volume.
	//
	// For example, setting the mountPath as /etc/secrets would mount the secret value files under the /etc/secrets directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount path: /etc/secrets
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudfunctions2_function#mount_path GoogleCloudfunctions2Function#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret.
	//
	// If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudfunctions2_function#project_id GoogleCloudfunctions2Function#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Name of the secret in secret manager (not the full resource name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudfunctions2_function#secret GoogleCloudfunctions2Function#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudfunctions2_function#versions GoogleCloudfunctions2Function#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
}

