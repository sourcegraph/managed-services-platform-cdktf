package computeroute


type ComputeRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_route#create ComputeRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_route#delete ComputeRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
