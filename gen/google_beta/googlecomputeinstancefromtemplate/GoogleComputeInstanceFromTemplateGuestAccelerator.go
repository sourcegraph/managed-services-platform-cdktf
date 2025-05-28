package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_instance_from_template#count GoogleComputeInstanceFromTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_compute_instance_from_template#type GoogleComputeInstanceFromTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

