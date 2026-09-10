package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsTlsDecrypt struct {
	// Specify whether to inspect encrypted HTTP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

