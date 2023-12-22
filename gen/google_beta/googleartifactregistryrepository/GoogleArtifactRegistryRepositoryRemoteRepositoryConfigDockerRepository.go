package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository struct {
	// Address of the remote repository. Default value: "DOCKER_HUB" Possible values: ["DOCKER_HUB"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

