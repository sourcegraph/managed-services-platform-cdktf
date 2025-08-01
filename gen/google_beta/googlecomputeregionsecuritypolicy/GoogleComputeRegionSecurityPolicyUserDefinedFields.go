package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyUserDefinedFields struct {
	// The base relative to which 'offset' is measured.
	//
	// Possible values are:
	// - IPV4: Points to the beginning of the IPv4 header.
	// - IPV6: Points to the beginning of the IPv6 header.
	// - TCP: Points to the beginning of the TCP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments.
	// - UDP: Points to the beginning of the UDP header, skipping over any IPv4 options or IPv6 extension headers. Not present for non-first fragments. Possible values: ["IPV4", "IPV6", "TCP", "UDP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#base GoogleComputeRegionSecurityPolicy#base}
	Base *string `field:"required" json:"base" yaml:"base"`
	// If specified, apply this mask (bitwise AND) to the field to ignore bits before matching.
	//
	// Encoded as a hexadecimal number (starting with "0x").
	// The last byte of the field (in network byte order) corresponds to the least significant byte of the mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#mask GoogleComputeRegionSecurityPolicy#mask}
	Mask *string `field:"optional" json:"mask" yaml:"mask"`
	// The name of this field. Must be unique within the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#name GoogleComputeRegionSecurityPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Offset of the first byte of the field (in network byte order) relative to 'base'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#offset GoogleComputeRegionSecurityPolicy#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Size of the field in bytes. Valid values: 1-4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#size GoogleComputeRegionSecurityPolicy#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

