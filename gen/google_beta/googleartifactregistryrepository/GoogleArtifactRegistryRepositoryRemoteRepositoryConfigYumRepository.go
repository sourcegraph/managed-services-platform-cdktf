package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository struct {
	// public_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

