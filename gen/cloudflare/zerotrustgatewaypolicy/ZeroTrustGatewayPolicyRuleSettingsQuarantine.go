package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsQuarantine struct {
	// Specify the types of files to sandbox.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#file_types ZeroTrustGatewayPolicy#file_types}
	FileTypes *[]*string `field:"optional" json:"fileTypes" yaml:"fileTypes"`
}

