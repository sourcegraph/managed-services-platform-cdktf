package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesMatchConfig struct {
	// CIDR IP address range. Maximum number of srcIpRanges allowed is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_security_policy#src_ip_ranges ComputeRegionSecurityPolicy#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

