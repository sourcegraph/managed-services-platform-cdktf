package zerotrustgatewaylogging


type ZeroTrustGatewayLoggingSettingsByRuleType struct {
	// Configure logging settings for DNS firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#dns ZeroTrustGatewayLogging#dns}
	Dns *ZeroTrustGatewayLoggingSettingsByRuleTypeDns `field:"optional" json:"dns" yaml:"dns"`
	// Configure logging settings for HTTP/HTTPS firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#http ZeroTrustGatewayLogging#http}
	Http *ZeroTrustGatewayLoggingSettingsByRuleTypeHttp `field:"optional" json:"http" yaml:"http"`
	// Configure logging settings for Network firewall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#l4 ZeroTrustGatewayLogging#l4}
	L4 *ZeroTrustGatewayLoggingSettingsByRuleTypeL4 `field:"optional" json:"l4" yaml:"l4"`
}

