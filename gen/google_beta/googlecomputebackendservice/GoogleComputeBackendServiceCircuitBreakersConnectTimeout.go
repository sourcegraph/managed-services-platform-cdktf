package googlecomputebackendservice


type GoogleComputeBackendServiceCircuitBreakersConnectTimeout struct {
	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_backend_service#seconds GoogleComputeBackendService#seconds}
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
	// Span of time that's a fraction of a second at nanosecond resolution.
	//
	// Durations less than one second are represented
	// with a 0 seconds field and a positive nanos field. Must
	// be from 0 to 999,999,999 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_backend_service#nanos GoogleComputeBackendService#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
}

