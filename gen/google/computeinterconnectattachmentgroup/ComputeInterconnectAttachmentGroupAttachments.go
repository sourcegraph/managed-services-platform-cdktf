package computeinterconnectattachmentgroup


type ComputeInterconnectAttachmentGroupAttachments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_interconnect_attachment_group#name ComputeInterconnectAttachmentGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_interconnect_attachment_group#attachment ComputeInterconnectAttachmentGroup#attachment}.
	Attachment *string `field:"optional" json:"attachment" yaml:"attachment"`
}

