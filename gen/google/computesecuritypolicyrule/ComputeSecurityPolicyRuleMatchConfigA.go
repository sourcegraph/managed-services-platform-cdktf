package computesecuritypolicyrule


type ComputeSecurityPolicyRuleMatchConfigA struct {
	// CIDR IP address range. Maximum number of srcIpRanges allowed is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/compute_security_policy_rule#src_ip_ranges ComputeSecurityPolicyRuleA#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}
