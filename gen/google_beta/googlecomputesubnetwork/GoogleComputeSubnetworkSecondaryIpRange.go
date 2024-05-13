package googlecomputesubnetwork


type GoogleComputeSubnetworkSecondaryIpRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_compute_subnetwork#range_name GoogleComputeSubnetwork#range_name}.
	RangeName *string `field:"optional" json:"rangeName" yaml:"rangeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_compute_subnetwork#reserved_internal_range GoogleComputeSubnetwork#reserved_internal_range}.
	ReservedInternalRange *string `field:"optional" json:"reservedInternalRange" yaml:"reservedInternalRange"`
}

