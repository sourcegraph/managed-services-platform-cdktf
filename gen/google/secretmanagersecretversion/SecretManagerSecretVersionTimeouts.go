package secretmanagersecretversion


type SecretManagerSecretVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/secret_manager_secret_version#create SecretManagerSecretVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/secret_manager_secret_version#delete SecretManagerSecretVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
