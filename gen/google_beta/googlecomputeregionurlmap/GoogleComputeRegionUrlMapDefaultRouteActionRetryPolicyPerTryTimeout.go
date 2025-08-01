package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicyPerTryTimeout struct {
	// Span of time that's a fraction of a second at nanosecond resolution.
	//
	// Durations less than one second are
	// represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#nanos GoogleComputeRegionUrlMap#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Span of time at a resolution of a second.
	//
	// Must be from 0 to 315,576,000,000 inclusive.
	// Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#seconds GoogleComputeRegionUrlMap#seconds}
	Seconds *string `field:"optional" json:"seconds" yaml:"seconds"`
}

