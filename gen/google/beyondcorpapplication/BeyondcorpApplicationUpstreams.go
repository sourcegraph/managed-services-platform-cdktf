package beyondcorpapplication


type BeyondcorpApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_application#egress_policy BeyondcorpApplication#egress_policy}
	EgressPolicy *BeyondcorpApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_application#network BeyondcorpApplication#network}
	Network *BeyondcorpApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
}

