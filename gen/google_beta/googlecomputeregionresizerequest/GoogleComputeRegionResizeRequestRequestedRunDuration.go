package googlecomputeregionresizerequest


type GoogleComputeRegionResizeRequestRequestedRunDuration struct {
	// Span of time at a resolution of a second.
	//
	// Must be from 600 to 604800 inclusive. Note: minimum and maximum allowed range for requestedRunDuration is 10 minutes (600 seconds) and 7 days(604800 seconds) correspondingly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_region_resize_request#seconds GoogleComputeRegionResizeRequest#seconds}
	Seconds *string `field:"required" json:"seconds" yaml:"seconds"`
	// Span of time that's a fraction of a second at nanosecond resolution.
	//
	// Durations less than one second are represented with a 0 seconds field and a positive nanos field. Must be from 0 to 999,999,999 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_region_resize_request#nanos GoogleComputeRegionResizeRequest#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
}

