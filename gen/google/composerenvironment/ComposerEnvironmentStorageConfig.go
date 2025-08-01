package composerenvironment


type ComposerEnvironmentStorageConfig struct {
	// Optional. Name of an existing Cloud Storage bucket to be used by the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#bucket ComposerEnvironment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
}

