package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsPayloadLog struct {
	// Enable or disable DLP Payload Logging for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_policy#enabled ZeroTrustGatewayPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

