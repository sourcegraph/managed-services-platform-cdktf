package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsAuditSsh struct {
	// Log all SSH commands.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_policy#command_logging ZeroTrustGatewayPolicy#command_logging}
	CommandLogging interface{} `field:"required" json:"commandLogging" yaml:"commandLogging"`
}

