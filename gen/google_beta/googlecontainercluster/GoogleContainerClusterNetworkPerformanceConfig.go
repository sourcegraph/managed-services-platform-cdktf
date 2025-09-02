package googlecontainercluster


type GoogleContainerClusterNetworkPerformanceConfig struct {
	// Specifies the total network bandwidth tier for NodePools in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#total_egress_bandwidth_tier GoogleContainerCluster#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

