package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleMatch struct {
	// Available values: "windows", "mac", "linux", "android", "ios", "chromeos".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_posture_rule#platform ZeroTrustDevicePostureRule#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

