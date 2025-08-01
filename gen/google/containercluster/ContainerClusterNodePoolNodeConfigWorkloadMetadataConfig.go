package containercluster


type ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#mode ContainerCluster#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

