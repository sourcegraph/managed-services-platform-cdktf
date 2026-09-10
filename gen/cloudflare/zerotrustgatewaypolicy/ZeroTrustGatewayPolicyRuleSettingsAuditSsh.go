package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsAuditSsh struct {
	// Enable SSH command logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#command_logging ZeroTrustGatewayPolicy#command_logging}
	CommandLogging interface{} `field:"optional" json:"commandLogging" yaml:"commandLogging"`
}

