package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA struct {
	// A list of site keys to be used during the validation of reCAPTCHA action-tokens.
	//
	// The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_security_policy_rule#action_token_site_keys GoogleComputeSecurityPolicyRuleA#action_token_site_keys}
	ActionTokenSiteKeys *[]*string `field:"optional" json:"actionTokenSiteKeys" yaml:"actionTokenSiteKeys"`
	// A list of site keys to be used during the validation of reCAPTCHA session-tokens.
	//
	// The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_security_policy_rule#session_token_site_keys GoogleComputeSecurityPolicyRuleA#session_token_site_keys}
	SessionTokenSiteKeys *[]*string `field:"optional" json:"sessionTokenSiteKeys" yaml:"sessionTokenSiteKeys"`
}

