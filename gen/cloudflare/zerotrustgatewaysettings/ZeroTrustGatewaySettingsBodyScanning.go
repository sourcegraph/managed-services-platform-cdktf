package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsBodyScanning struct {
	// Body scanning inspection mode. Available values: `deep`, `shallow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_settings#inspection_mode ZeroTrustGatewaySettings#inspection_mode}
	InspectionMode *string `field:"required" json:"inspectionMode" yaml:"inspectionMode"`
}

