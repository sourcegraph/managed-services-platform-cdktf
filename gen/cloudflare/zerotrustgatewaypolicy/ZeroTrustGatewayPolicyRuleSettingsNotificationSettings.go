package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsNotificationSettings struct {
	// Enable notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#enabled ZeroTrustGatewayPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates whether to pass the context information as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#include_context ZeroTrustGatewayPolicy#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// Customize the message shown in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#msg ZeroTrustGatewayPolicy#msg}
	Msg *string `field:"optional" json:"msg" yaml:"msg"`
	// Defines an optional URL to direct users to additional information. If unset, the notification opens a block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#support_url ZeroTrustGatewayPolicy#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

