package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryDockerConfig struct {
	// The repository which enabled this flag prevents all tags from being modified, moved or deleted.
	//
	// This does not prevent tags from being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_artifact_registry_repository#immutable_tags GoogleArtifactRegistryRepository#immutable_tags}
	ImmutableTags interface{} `field:"optional" json:"immutableTags" yaml:"immutableTags"`
}

