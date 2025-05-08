package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance_from_machine_image#count GoogleComputeInstanceFromMachineImage#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance_from_machine_image#type GoogleComputeInstanceFromMachineImage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

