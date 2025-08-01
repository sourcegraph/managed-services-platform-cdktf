package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsL4Override struct {
	// IPv4 or IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#ip ZeroTrustGatewayPolicy#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// A port number to use for TCP/UDP overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#port ZeroTrustGatewayPolicy#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

