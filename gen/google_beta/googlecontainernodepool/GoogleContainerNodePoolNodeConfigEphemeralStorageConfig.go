package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigEphemeralStorageConfig struct {
	// Number of local SSDs to use to back ephemeral storage.
	//
	// Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool#local_ssd_count GoogleContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

