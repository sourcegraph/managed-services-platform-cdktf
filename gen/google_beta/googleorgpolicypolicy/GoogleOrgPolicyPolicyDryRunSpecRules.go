package googleorgpolicypolicy


type GoogleOrgPolicyPolicyDryRunSpecRules struct {
	// Setting this to `"TRUE"` means that all values are allowed.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_org_policy_policy#allow_all GoogleOrgPolicyPolicy#allow_all}
	AllowAll *string `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_org_policy_policy#condition GoogleOrgPolicyPolicy#condition}
	Condition *GoogleOrgPolicyPolicyDryRunSpecRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to `"TRUE"` means that all values are denied.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_org_policy_policy#deny_all GoogleOrgPolicyPolicy#deny_all}
	DenyAll *string `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If `"TRUE"`, then the policy is enforced.
	//
	// If `"FALSE"`, then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_org_policy_policy#enforce GoogleOrgPolicyPolicy#enforce}
	Enforce *string `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_org_policy_policy#values GoogleOrgPolicyPolicy#values}
	Values *GoogleOrgPolicyPolicyDryRunSpecRulesValues `field:"optional" json:"values" yaml:"values"`
}

