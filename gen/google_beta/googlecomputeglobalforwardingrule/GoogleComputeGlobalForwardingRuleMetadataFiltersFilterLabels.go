package googlecomputeglobalforwardingrule


type GoogleComputeGlobalForwardingRuleMetadataFiltersFilterLabels struct {
	// Name of the metadata label. The length must be between 1 and 1024 characters, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_global_forwarding_rule#name GoogleComputeGlobalForwardingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value that the label must match. The value has a maximum length of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_global_forwarding_rule#value GoogleComputeGlobalForwardingRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

