package computepacketmirroring


type ComputePacketMirroringNetwork struct {
	// The full self_link URL of the network where this rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_packet_mirroring#url ComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

