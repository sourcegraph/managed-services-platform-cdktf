package containercluster


type ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig struct {
	// Name for pod secondary ipv4 range which has the actual range defined ahead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#pod_range_names ContainerCluster#pod_range_names}
	PodRangeNames *[]*string `field:"required" json:"podRangeNames" yaml:"podRangeNames"`
}

