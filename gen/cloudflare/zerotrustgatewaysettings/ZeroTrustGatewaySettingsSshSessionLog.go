package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSshSessionLog struct {
	// Public key used to encrypt ssh session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_settings#public_key ZeroTrustGatewaySettings#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

