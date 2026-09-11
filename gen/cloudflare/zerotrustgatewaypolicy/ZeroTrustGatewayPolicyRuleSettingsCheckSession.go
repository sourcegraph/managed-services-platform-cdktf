package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsCheckSession struct {
	// Sets the required session freshness threshold. The API returns a normalized version of this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Enable session enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#enforce ZeroTrustGatewayPolicy#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
}

