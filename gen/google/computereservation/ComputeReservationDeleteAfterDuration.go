package computereservation


type ComputeReservationDeleteAfterDuration struct {
	// Number of nanoseconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_reservation#nanos ComputeReservation#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Number of seconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_reservation#seconds ComputeReservation#seconds}
	Seconds *string `field:"optional" json:"seconds" yaml:"seconds"`
}

