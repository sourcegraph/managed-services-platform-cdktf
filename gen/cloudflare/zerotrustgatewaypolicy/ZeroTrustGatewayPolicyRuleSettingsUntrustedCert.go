package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsUntrustedCert struct {
	// Defines the action performed when an untrusted certificate seen.
	//
	// The default action an error with HTTP code 526.
	// Available values: "pass_through", "block", "error".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#action ZeroTrustGatewayPolicy#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

