package googlecomputewiregroup


type GoogleComputeWireGroupWireProperties struct {
	// The unmetered bandwidth setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_wire_group#bandwidth_unmetered GoogleComputeWireGroup#bandwidth_unmetered}
	BandwidthUnmetered *float64 `field:"optional" json:"bandwidthUnmetered" yaml:"bandwidthUnmetered"`
	// Response when a fault is detected in a pseudowire: NONE: default.
	//
	// DISABLE_PORT: set the port line protocol down when inline probes detect a fault. This setting is only permitted on port mode pseudowires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_wire_group#fault_response GoogleComputeWireGroup#fault_response}
	FaultResponse *string `field:"optional" json:"faultResponse" yaml:"faultResponse"`
}

