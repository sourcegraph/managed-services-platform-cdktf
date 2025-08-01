package zerotrustgatewaylogging


type ZeroTrustGatewayLoggingSettingsByRuleType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_logging#dns ZeroTrustGatewayLogging#dns}.
	Dns *ZeroTrustGatewayLoggingSettingsByRuleTypeDns `field:"optional" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_logging#http ZeroTrustGatewayLogging#http}.
	Http *ZeroTrustGatewayLoggingSettingsByRuleTypeHttp `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_logging#l4 ZeroTrustGatewayLogging#l4}.
	L4 *ZeroTrustGatewayLoggingSettingsByRuleTypeL4 `field:"optional" json:"l4" yaml:"l4"`
}

