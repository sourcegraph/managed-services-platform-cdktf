package zerotrustgatewaylogging


type ZeroTrustGatewayLoggingSettingsByRuleTypeL4 struct {
	// Log all requests to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_logging#log_all ZeroTrustGatewayLogging#log_all}
	LogAll interface{} `field:"optional" json:"logAll" yaml:"logAll"`
	// Log only blocking requests to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_logging#log_blocks ZeroTrustGatewayLogging#log_blocks}
	LogBlocks interface{} `field:"optional" json:"logBlocks" yaml:"logBlocks"`
}

