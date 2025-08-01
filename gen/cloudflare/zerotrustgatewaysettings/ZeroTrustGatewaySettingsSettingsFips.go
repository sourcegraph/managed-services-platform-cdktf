package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_settings#tls ZeroTrustGatewaySettings#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

