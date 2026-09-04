package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings struct {
	// Specify whether to enable notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify whether to include context information as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#include_context ZeroTrustGatewaySettings#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// Specify the message to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#msg ZeroTrustGatewaySettings#msg}
	Msg *string `field:"optional" json:"msg" yaml:"msg"`
	// Specify a URL that directs users to more information. If unset, the notification opens a block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#support_url ZeroTrustGatewaySettings#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

