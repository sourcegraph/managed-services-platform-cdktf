package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsExtendedEmailMatching struct {
	// Specify whether to match all variants of user emails (with + or .
	//
	// modifiers) used as criteria in Firewall policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

