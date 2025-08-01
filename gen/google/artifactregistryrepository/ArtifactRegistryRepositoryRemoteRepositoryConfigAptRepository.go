package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository struct {
	// public_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

