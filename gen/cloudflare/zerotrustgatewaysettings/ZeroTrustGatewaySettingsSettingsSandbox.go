package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsSandbox struct {
	// Specify whether to enable the sandbox.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the action to take when the system cannot scan the file. Available values: "allow", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#fallback_action ZeroTrustGatewaySettings#fallback_action}
	FallbackAction *string `field:"optional" json:"fallbackAction" yaml:"fallbackAction"`
}

