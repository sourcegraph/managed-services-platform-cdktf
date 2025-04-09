package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://pypi.io"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_artifact_registry_repository#uri GoogleArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

