package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls struct {
	// Configure copy behavior.
	//
	// If set to remote_only, users cannot copy isolated content from the remote browser to the local clipboard. If this field is absent, copying remains enabled. Applies only when version == "v2".
	// Available values: "enabled", "disabled", "remote_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#copy ZeroTrustGatewayPolicy#copy}
	Copy *string `field:"optional" json:"copy" yaml:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#dcp ZeroTrustGatewayPolicy#dcp}
	Dcp interface{} `field:"optional" json:"dcp" yaml:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#dd ZeroTrustGatewayPolicy#dd}
	Dd interface{} `field:"optional" json:"dd" yaml:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#dk ZeroTrustGatewayPolicy#dk}
	Dk interface{} `field:"optional" json:"dk" yaml:"dk"`
	// Configure download behavior.
	//
	// When set to remote_only, users can view downloads but cannot save them. Applies only when version == "v2".
	// Available values: "enabled", "disabled", "remote_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#download ZeroTrustGatewayPolicy#download}
	Download *string `field:"optional" json:"download" yaml:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#dp ZeroTrustGatewayPolicy#dp}
	Dp interface{} `field:"optional" json:"dp" yaml:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#du ZeroTrustGatewayPolicy#du}
	Du interface{} `field:"optional" json:"du" yaml:"du"`
	// Configure keyboard usage behavior.
	//
	// If this field is absent, keyboard usage remains enabled. Applies only when version == "v2".
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#keyboard ZeroTrustGatewayPolicy#keyboard}
	Keyboard *string `field:"optional" json:"keyboard" yaml:"keyboard"`
	// Configure paste behavior.
	//
	// If set to remote_only, users cannot paste content from the local clipboard into isolated pages. If this field is absent, pasting remains enabled. Applies only when version == "v2".
	// Available values: "enabled", "disabled", "remote_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#paste ZeroTrustGatewayPolicy#paste}
	Paste *string `field:"optional" json:"paste" yaml:"paste"`
	// Configure print behavior. Default, Printing is enabled. Applies only when version == "v2". Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#printing ZeroTrustGatewayPolicy#printing}
	Printing *string `field:"optional" json:"printing" yaml:"printing"`
	// Configure upload behavior.
	//
	// If this field is absent, uploading remains enabled. Applies only when version == "v2".
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#upload ZeroTrustGatewayPolicy#upload}
	Upload *string `field:"optional" json:"upload" yaml:"upload"`
	// Indicate which version of the browser isolation controls should apply. Available values: "v1", "v2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#version ZeroTrustGatewayPolicy#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

