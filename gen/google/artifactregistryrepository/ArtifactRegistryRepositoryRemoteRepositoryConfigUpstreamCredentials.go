package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials struct {
	// username_password_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#username_password_credentials ArtifactRegistryRepository#username_password_credentials}
	UsernamePasswordCredentials *ArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials `field:"optional" json:"usernamePasswordCredentials" yaml:"usernamePasswordCredentials"`
}

