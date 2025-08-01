package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen.
	//
	// The default action is an error with HTTP code 526.
	// Available values: "pass_through", "block", "error".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#action ZeroTrustGatewayPolicy#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

