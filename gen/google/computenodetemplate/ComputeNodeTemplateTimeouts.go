package computenodetemplate


type ComputeNodeTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_node_template#create ComputeNodeTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_node_template#delete ComputeNodeTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

