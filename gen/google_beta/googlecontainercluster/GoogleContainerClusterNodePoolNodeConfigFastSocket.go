package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigFastSocket struct {
	// Whether or not NCCL Fast Socket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

