package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsExtendedEmailMatching struct {
	// Specify whether to match all variants of user emails (with + or .
	//
	// modifiers) used as criteria in Firewall policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicate that this setting was shared via the Orgs API and read only for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#read_only ZeroTrustGatewaySettings#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Indicate the account tag of the account that shared this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#source_account ZeroTrustGatewaySettings#source_account}
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// Indicate the version number of the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#version ZeroTrustGatewaySettings#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

