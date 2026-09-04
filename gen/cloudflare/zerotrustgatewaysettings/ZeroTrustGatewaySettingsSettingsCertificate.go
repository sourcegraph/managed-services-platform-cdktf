package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsCertificate struct {
	// Specify the UUID of the certificate used for interception.
	//
	// Ensure the certificate is available at the edge(previously called 'active'). A nil UUID directs Cloudflare to use the Root CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#id ZeroTrustGatewaySettings#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

