package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleRedirectOptionsA struct {
	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_security_policy_rule#target GoogleComputeSecurityPolicyRuleA#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Type of the redirect action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_security_policy_rule#type GoogleComputeSecurityPolicyRuleA#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

