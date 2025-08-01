package computenodetemplate


type ComputeNodeTemplateAccelerators struct {
	// The number of the guest accelerator cards exposed to this node template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_node_template#accelerator_count ComputeNodeTemplate#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Full or partial URL of the accelerator type resource to expose to this node template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_node_template#accelerator_type ComputeNodeTemplate#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
}

