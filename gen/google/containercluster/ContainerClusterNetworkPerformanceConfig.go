package containercluster


type ContainerClusterNetworkPerformanceConfig struct {
	// Specifies the total network bandwidth tier for NodePools in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#total_egress_bandwidth_tier ContainerCluster#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

