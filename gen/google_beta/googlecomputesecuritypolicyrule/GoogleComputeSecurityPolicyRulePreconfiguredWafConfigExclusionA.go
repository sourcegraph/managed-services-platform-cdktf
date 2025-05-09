package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRulePreconfiguredWafConfigExclusionA struct {
	// Target WAF rule set to apply the preconfigured WAF exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#target_rule_set GoogleComputeSecurityPolicyRuleA#target_rule_set}
	TargetRuleSet *string `field:"required" json:"targetRuleSet" yaml:"targetRuleSet"`
	// request_cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#request_cookie GoogleComputeSecurityPolicyRuleA#request_cookie}
	RequestCookie interface{} `field:"optional" json:"requestCookie" yaml:"requestCookie"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#request_header GoogleComputeSecurityPolicyRuleA#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// request_query_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#request_query_param GoogleComputeSecurityPolicyRuleA#request_query_param}
	RequestQueryParam interface{} `field:"optional" json:"requestQueryParam" yaml:"requestQueryParam"`
	// request_uri block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#request_uri GoogleComputeSecurityPolicyRuleA#request_uri}
	RequestUri interface{} `field:"optional" json:"requestUri" yaml:"requestUri"`
	// A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion.
	//
	// If omitted, it refers to all the rule IDs under the WAF rule set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_security_policy_rule#target_rule_ids GoogleComputeSecurityPolicyRuleA#target_rule_ids}
	TargetRuleIds *[]*string `field:"optional" json:"targetRuleIds" yaml:"targetRuleIds"`
}

