package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://repo.maven.apache.org/maven2"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#uri ArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

