package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Set to true to enable session enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#enforce ZeroTrustGatewayPolicy#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
}

