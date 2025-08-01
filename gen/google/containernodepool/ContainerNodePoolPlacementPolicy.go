package containernodepool


type ContainerNodePoolPlacementPolicy struct {
	// Type defines the type of placement policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_node_pool#type ContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If set, refers to the name of a custom resource policy supplied by the user.
	//
	// The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_node_pool#policy_name ContainerNodePool#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The TPU topology like "2x4" or "2x2x2". https://cloud.google.com/kubernetes-engine/docs/concepts/plan-tpus#topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_node_pool#tpu_topology ContainerNodePool#tpu_topology}
	TpuTopology *string `field:"optional" json:"tpuTopology" yaml:"tpuTopology"`
}

