package computeattacheddisk


type ComputeAttachedDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_attached_disk#create ComputeAttachedDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_attached_disk#delete ComputeAttachedDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

