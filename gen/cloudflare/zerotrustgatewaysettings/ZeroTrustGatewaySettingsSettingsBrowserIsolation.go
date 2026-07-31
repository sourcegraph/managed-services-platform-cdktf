package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBrowserIsolation struct {
	// Specify whether to enable non-identity onramp support for Browser Isolation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#non_identity_enabled ZeroTrustGatewaySettings#non_identity_enabled}
	NonIdentityEnabled interface{} `field:"optional" json:"nonIdentityEnabled" yaml:"nonIdentityEnabled"`
	// Specify whether to enable Clientless Browser Isolation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#url_browser_isolation_enabled ZeroTrustGatewaySettings#url_browser_isolation_enabled}
	UrlBrowserIsolationEnabled interface{} `field:"optional" json:"urlBrowserIsolationEnabled" yaml:"urlBrowserIsolationEnabled"`
}

