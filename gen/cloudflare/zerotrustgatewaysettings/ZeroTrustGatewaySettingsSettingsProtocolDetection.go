package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsProtocolDetection struct {
	// Specify whether to detect protocols from the initial bytes of client traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

