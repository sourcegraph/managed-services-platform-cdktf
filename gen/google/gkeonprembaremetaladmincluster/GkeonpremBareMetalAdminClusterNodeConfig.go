package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterNodeConfig struct {
	// The maximum number of pods a node can run.
	//
	// The size of the CIDR range
	// assigned to the node will be derived from this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#max_pods_per_node GkeonpremBareMetalAdminCluster#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

