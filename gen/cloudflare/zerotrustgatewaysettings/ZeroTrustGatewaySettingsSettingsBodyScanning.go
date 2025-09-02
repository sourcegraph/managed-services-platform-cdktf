package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_settings#inspection_mode ZeroTrustGatewaySettings#inspection_mode}
	InspectionMode *string `field:"optional" json:"inspectionMode" yaml:"inspectionMode"`
}

