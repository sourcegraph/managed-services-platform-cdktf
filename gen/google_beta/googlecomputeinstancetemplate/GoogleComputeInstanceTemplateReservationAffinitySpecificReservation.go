package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateReservationAffinitySpecificReservation struct {
	// Corresponds to the label key of a reservation resource.
	//
	// To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_compute_instance_template#key GoogleComputeInstanceTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Corresponds to the label values of a reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_compute_instance_template#values GoogleComputeInstanceTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

