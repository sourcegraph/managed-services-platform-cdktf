package computerouterpeer


type ComputeRouterPeerCustomLearnedIpRanges struct {
	// The IP range to learn. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_router_peer#range ComputeRouterPeer#range}
	Range *string `field:"required" json:"range" yaml:"range"`
}

