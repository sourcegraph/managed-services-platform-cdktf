package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyRules struct {
	// The Action to perform when the rule is matched. The following are the valid actions:.
	//
	// * allow: allow access to target.
	//
	// * deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for STATUS are 403, 404, and 502.
	//
	// * rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rateLimitOptions to be set.
	//
	// * redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR.
	//
	// * throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rateLimitOptions to be set for this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#action GoogleComputeRegionSecurityPolicy#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a positive value between 0 and 2147483647.
	// Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#priority GoogleComputeRegionSecurityPolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#description GoogleComputeRegionSecurityPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#match GoogleComputeRegionSecurityPolicy#match}
	Match *GoogleComputeRegionSecurityPolicyRulesMatch `field:"optional" json:"match" yaml:"match"`
	// network_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#network_match GoogleComputeRegionSecurityPolicy#network_match}
	NetworkMatch *GoogleComputeRegionSecurityPolicyRulesNetworkMatch `field:"optional" json:"networkMatch" yaml:"networkMatch"`
	// preconfigured_waf_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#preconfigured_waf_config GoogleComputeRegionSecurityPolicy#preconfigured_waf_config}
	PreconfiguredWafConfig *GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfig `field:"optional" json:"preconfiguredWafConfig" yaml:"preconfiguredWafConfig"`
	// If set to true, the specified action is not enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#preview GoogleComputeRegionSecurityPolicy#preview}
	Preview interface{} `field:"optional" json:"preview" yaml:"preview"`
	// rate_limit_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_security_policy#rate_limit_options GoogleComputeRegionSecurityPolicy#rate_limit_options}
	RateLimitOptions *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions `field:"optional" json:"rateLimitOptions" yaml:"rateLimitOptions"`
}

