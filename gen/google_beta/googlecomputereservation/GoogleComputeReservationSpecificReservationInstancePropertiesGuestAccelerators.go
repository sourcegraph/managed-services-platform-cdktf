package googlecomputereservation


type GoogleComputeReservationSpecificReservationInstancePropertiesGuestAccelerators struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_reservation#accelerator_count GoogleComputeReservation#accelerator_count}
	AcceleratorCount *float64 `field:"required" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The full or partial URL of the accelerator type to attach to this instance. For example: 'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100'.
	//
	// If you are creating an instance template, specify only the accelerator name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_reservation#accelerator_type GoogleComputeReservation#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
}

