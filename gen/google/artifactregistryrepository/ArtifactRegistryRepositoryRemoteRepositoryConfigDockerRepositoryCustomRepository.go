package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://registry-1.docker.io"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#uri ArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

