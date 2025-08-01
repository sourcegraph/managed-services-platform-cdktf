package computesecuritypolicy


type ComputeSecurityPolicyRulePreconfiguredWafConfig struct {
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy#exclusion ComputeSecurityPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}

