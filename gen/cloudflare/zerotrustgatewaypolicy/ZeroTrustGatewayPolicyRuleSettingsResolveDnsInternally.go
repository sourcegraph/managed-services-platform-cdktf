package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally struct {
	// Specify the fallback behavior to apply when the internal DNS response code differs from 'NOERROR' or when the response data contains only CNAME records for 'A' or 'AAAA' queries.
	//
	// Available values: "none", "public_dns".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#fallback ZeroTrustGatewayPolicy#fallback}
	Fallback *string `field:"optional" json:"fallback" yaml:"fallback"`
	// Specify the internal DNS view identifier to pass to the internal DNS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#view_id ZeroTrustGatewayPolicy#view_id}
	ViewId *string `field:"optional" json:"viewId" yaml:"viewId"`
}

