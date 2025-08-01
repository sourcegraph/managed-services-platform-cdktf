package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsHostSelector struct {
	// Enable filtering via hosts for egress policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

