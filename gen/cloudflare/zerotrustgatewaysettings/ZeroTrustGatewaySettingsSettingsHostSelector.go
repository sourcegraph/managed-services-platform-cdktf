package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsHostSelector struct {
	// Specify whether to enable filtering via hosts for egress policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

