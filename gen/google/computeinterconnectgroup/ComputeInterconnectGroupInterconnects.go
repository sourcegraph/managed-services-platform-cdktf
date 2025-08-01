package computeinterconnectgroup


type ComputeInterconnectGroupInterconnects struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_interconnect_group#name ComputeInterconnectGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of an Interconnect in this group. All Interconnects in the group are unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_interconnect_group#interconnect ComputeInterconnectGroup#interconnect}
	Interconnect *string `field:"optional" json:"interconnect" yaml:"interconnect"`
}

