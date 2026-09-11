package zerotrustaccessgroup


type ZeroTrustAccessGroupIncludeUserRiskScore struct {
	// A list of risk score levels to match. Values can be low, medium, high, or unscored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_group#user_risk_score ZeroTrustAccessGroup#user_risk_score}
	UserRiskScore *[]*string `field:"required" json:"userRiskScore" yaml:"userRiskScore"`
}

