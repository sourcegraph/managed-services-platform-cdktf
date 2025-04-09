package googlecomputeinstance


type GoogleComputeInstanceSchedulingGracefulShutdownMaxDuration struct {
	// Span of time at a resolution of a second.
	//
	// The value must be between 1 and 3600, which is 3,600 seconds (one hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_instance#seconds GoogleComputeInstance#seconds}
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
	// Span of time that's a fraction of a second at nanosecond 													resolution.
	//
	// Durations less than one second are represented
	// 													with a 0 seconds field and a positive nanos field. Must
	// 													be from 0 to 999,999,999 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_instance#nanos GoogleComputeInstance#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
}

