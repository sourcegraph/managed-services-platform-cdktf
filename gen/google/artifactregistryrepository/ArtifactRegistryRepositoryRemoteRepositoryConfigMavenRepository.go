package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository struct {
	// custom_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#custom_repository ArtifactRegistryRepository#custom_repository}
	CustomRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryCustomRepository `field:"optional" json:"customRepository" yaml:"customRepository"`
	// Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

