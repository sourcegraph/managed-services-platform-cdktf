package orgpolicypolicy


type OrgPolicyPolicyDryRunSpecRules struct {
	// Setting this to true means that all values are allowed.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/org_policy_policy#allow_all OrgPolicyPolicy#allow_all}
	AllowAll *string `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/org_policy_policy#condition OrgPolicyPolicy#condition}
	Condition *OrgPolicyPolicyDryRunSpecRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to true means that all values are denied.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/org_policy_policy#deny_all OrgPolicyPolicy#deny_all}
	DenyAll *string `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If `true`, then the policy is enforced.
	//
	// If `false`, then any configuration is acceptable. This field can be set only in policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/org_policy_policy#enforce OrgPolicyPolicy#enforce}
	Enforce *string `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/org_policy_policy#values OrgPolicyPolicy#values}
	Values *OrgPolicyPolicyDryRunSpecRulesValues `field:"optional" json:"values" yaml:"values"`
}

