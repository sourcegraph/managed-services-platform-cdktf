package containercluster


type ContainerClusterEnterpriseConfig struct {
	// Indicates the desired cluster tier. Available options include STANDARD and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/container_cluster#desired_tier ContainerCluster#desired_tier}
	DesiredTier *string `field:"optional" json:"desiredTier" yaml:"desiredTier"`
}

