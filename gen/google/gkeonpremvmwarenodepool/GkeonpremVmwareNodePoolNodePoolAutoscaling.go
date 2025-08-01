package gkeonpremvmwarenodepool


type GkeonpremVmwareNodePoolNodePoolAutoscaling struct {
	// Maximum number of replicas in the NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_node_pool#max_replicas GkeonpremVmwareNodePool#max_replicas}
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// Minimum number of replicas in the NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_node_pool#min_replicas GkeonpremVmwareNodePool#min_replicas}
	MinReplicas *float64 `field:"required" json:"minReplicas" yaml:"minReplicas"`
}

