package computesecuritypolicyrule


type ComputeSecurityPolicyRuleMatchA struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#config ComputeSecurityPolicyRuleA#config}
	Config *ComputeSecurityPolicyRuleMatchConfigA `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#expr ComputeSecurityPolicyRuleA#expr}
	Expr *ComputeSecurityPolicyRuleMatchExprA `field:"optional" json:"expr" yaml:"expr"`
	// expr_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#expr_options ComputeSecurityPolicyRuleA#expr_options}
	ExprOptions *ComputeSecurityPolicyRuleMatchExprOptionsA `field:"optional" json:"exprOptions" yaml:"exprOptions"`
	// Preconfigured versioned expression.
	//
	// If this field is specified, config must also be specified.
	// Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#versioned_expr ComputeSecurityPolicyRuleA#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

