package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsUntrustedCert struct {
	// Action to be taken when the SSL certificate of upstream is invalid. Available values: `pass_through`, `block`, `error`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_policy#action ZeroTrustGatewayPolicy#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

