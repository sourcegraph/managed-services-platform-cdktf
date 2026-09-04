package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsPayloadLog struct {
	// Enable DLP payload logging for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#enabled ZeroTrustGatewayPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

