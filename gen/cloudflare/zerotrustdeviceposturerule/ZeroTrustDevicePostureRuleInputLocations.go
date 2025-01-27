package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleInputLocations struct {
	// List of paths to check for client certificate rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_rule#paths ZeroTrustDevicePostureRule#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// List of trust stores to check for client certificate rule. Available values: `system`, `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_rule#trust_stores ZeroTrustDevicePostureRule#trust_stores}
	TrustStores *[]*string `field:"optional" json:"trustStores" yaml:"trustStores"`
}

