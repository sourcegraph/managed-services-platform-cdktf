package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigCommonRepository struct {
	// One of: a.
	//
	// Artifact Registry Repository resource, e.g. 'projects/UPSTREAM_PROJECT_ID/locations/REGION/repositories/UPSTREAM_REPOSITORY'
	// b. URI to the registry, e.g. '"https://registry-1.docker.io"'
	// c. URI to Artifact Registry Repository, e.g. '"https://REGION-docker.pkg.dev/UPSTREAM_PROJECT_ID/UPSTREAM_REPOSITORY"'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/artifact_registry_repository#uri ArtifactRegistryRepository#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

