package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileThreatPreventionProfile struct {
	// antivirus_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_security_security_profile#antivirus_overrides GoogleNetworkSecuritySecurityProfile#antivirus_overrides}
	AntivirusOverrides interface{} `field:"optional" json:"antivirusOverrides" yaml:"antivirusOverrides"`
	// severity_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_security_security_profile#severity_overrides GoogleNetworkSecuritySecurityProfile#severity_overrides}
	SeverityOverrides interface{} `field:"optional" json:"severityOverrides" yaml:"severityOverrides"`
	// threat_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_security_security_profile#threat_overrides GoogleNetworkSecuritySecurityProfile#threat_overrides}
	ThreatOverrides interface{} `field:"optional" json:"threatOverrides" yaml:"threatOverrides"`
}

