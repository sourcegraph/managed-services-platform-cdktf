package googlecontainernodepool


type GoogleContainerNodePoolAutoscaling struct {
	// Location policy specifies the algorithm used when scaling-up the node pool.
	//
	// "BALANCED" - Is a best effort policy that aims to balance the sizes of available zones. "ANY" - Instructs the cluster autoscaler to prioritize utilization of unused reservations, and reduces preemption risk for Spot VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#location_policy GoogleContainerNodePool#location_policy}
	LocationPolicy *string `field:"optional" json:"locationPolicy" yaml:"locationPolicy"`
	// Maximum number of nodes per zone in the node pool.
	//
	// Must be >= min_node_count. Cannot be used with total limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#max_node_count GoogleContainerNodePool#max_node_count}
	MaxNodeCount *float64 `field:"optional" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes per zone in the node pool.
	//
	// Must be >=0 and <= max_node_count. Cannot be used with total limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#min_node_count GoogleContainerNodePool#min_node_count}
	MinNodeCount *float64 `field:"optional" json:"minNodeCount" yaml:"minNodeCount"`
	// Maximum number of all nodes in the node pool.
	//
	// Must be >= total_min_node_count. Cannot be used with per zone limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#total_max_node_count GoogleContainerNodePool#total_max_node_count}
	TotalMaxNodeCount *float64 `field:"optional" json:"totalMaxNodeCount" yaml:"totalMaxNodeCount"`
	// Minimum number of all nodes in the node pool.
	//
	// Must be >=0 and <= total_max_node_count. Cannot be used with per zone limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#total_min_node_count GoogleContainerNodePool#total_min_node_count}
	TotalMinNodeCount *float64 `field:"optional" json:"totalMinNodeCount" yaml:"totalMinNodeCount"`
}

