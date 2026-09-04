package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsInspection struct {
	// Define the proxy inspection mode.
	//
	// 1. static: Gateway applies static inspection to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on TCP(443) and UDP(443).   2. dynamic: Gateway applies protocol detection to inspect HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect HTTPS traffic.
	// Available values: "static", "dynamic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#mode ZeroTrustGatewaySettings#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

