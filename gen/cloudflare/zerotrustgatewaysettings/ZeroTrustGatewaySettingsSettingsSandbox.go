package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsSandbox struct {
	// Enable sandbox.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Action to take when the file cannot be scanned. Available values: "allow", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_settings#fallback_action ZeroTrustGatewaySettings#fallback_action}
	FallbackAction *string `field:"optional" json:"fallbackAction" yaml:"fallbackAction"`
}

