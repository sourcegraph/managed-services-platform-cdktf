package googlecomputerouternat


type GoogleComputeRouterNatSubnetwork struct {
	// Self-link of subnetwork to NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_nat#name GoogleComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of options for which source IPs in the subnetwork should have NAT enabled. Supported values include: 'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES', 'PRIMARY_IP_RANGE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_nat#source_ip_ranges_to_nat GoogleComputeRouterNat#source_ip_ranges_to_nat}
	SourceIpRangesToNat *[]*string `field:"required" json:"sourceIpRangesToNat" yaml:"sourceIpRangesToNat"`
	// List of the secondary ranges of the subnetwork that are allowed to use NAT.
	//
	// This can be populated only if
	// 'LIST_OF_SECONDARY_IP_RANGES' is one of the values in
	// sourceIpRangesToNat
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_nat#secondary_ip_range_names GoogleComputeRouterNat#secondary_ip_range_names}
	SecondaryIpRangeNames *[]*string `field:"optional" json:"secondaryIpRangeNames" yaml:"secondaryIpRangeNames"`
}

