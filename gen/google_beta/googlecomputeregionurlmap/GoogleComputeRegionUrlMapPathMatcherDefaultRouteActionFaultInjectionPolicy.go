package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#abort GoogleComputeRegionUrlMap#abort}
	Abort *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#delay GoogleComputeRegionUrlMap#delay}
	Delay *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

