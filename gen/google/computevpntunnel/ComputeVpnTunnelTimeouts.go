package computevpntunnel


type ComputeVpnTunnelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_vpn_tunnel#create ComputeVpnTunnel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_vpn_tunnel#delete ComputeVpnTunnel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_vpn_tunnel#update ComputeVpnTunnel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

