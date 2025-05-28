package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryMavenConfig struct {
	// The repository with this flag will allow publishing the same snapshot versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_artifact_registry_repository#allow_snapshot_overwrites GoogleArtifactRegistryRepository#allow_snapshot_overwrites}
	AllowSnapshotOverwrites interface{} `field:"optional" json:"allowSnapshotOverwrites" yaml:"allowSnapshotOverwrites"`
	// Version policy defines the versions that the registry will accept. Default value: "VERSION_POLICY_UNSPECIFIED" Possible values: ["VERSION_POLICY_UNSPECIFIED", "RELEASE", "SNAPSHOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_artifact_registry_repository#version_policy GoogleArtifactRegistryRepository#version_policy}
	VersionPolicy *string `field:"optional" json:"versionPolicy" yaml:"versionPolicy"`
}

