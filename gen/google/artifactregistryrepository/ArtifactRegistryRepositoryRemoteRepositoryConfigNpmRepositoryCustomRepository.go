package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://registry.npmjs.org"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/artifact_registry_repository#uri ArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

