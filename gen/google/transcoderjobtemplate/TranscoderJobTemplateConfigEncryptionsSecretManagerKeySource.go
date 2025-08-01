package transcoderjobtemplate


type TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource struct {
	// The name of the Secret Version containing the encryption key in the following format: projects/{project}/secrets/{secret_id}/versions/{version_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/transcoder_job_template#secret_version TranscoderJobTemplate#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

