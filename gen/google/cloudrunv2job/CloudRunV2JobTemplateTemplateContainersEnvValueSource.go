package cloudrunv2job


type CloudRunV2JobTemplateTemplateContainersEnvValueSource struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_v2_job#secret_key_ref CloudRunV2Job#secret_key_ref}
	SecretKeyRef *CloudRunV2JobTemplateTemplateContainersEnvValueSourceSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}
