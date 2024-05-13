package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository struct {
	// custom_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_artifact_registry_repository#custom_repository GoogleArtifactRegistryRepository#custom_repository}
	CustomRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository `field:"optional" json:"customRepository" yaml:"customRepository"`
	// Address of the remote repository. Default value: "PYPI" Possible values: ["PYPI"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

