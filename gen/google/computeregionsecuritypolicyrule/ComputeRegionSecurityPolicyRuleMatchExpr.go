package computeregionsecuritypolicyrule


type ComputeRegionSecurityPolicyRuleMatchExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_security_policy_rule#expression ComputeRegionSecurityPolicyRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

