package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsCustomCertificate struct {
	// Specify whether to enable a custom certificate authority for signing Gateway traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Indicate the internal certificate status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#binding_status ZeroTrustGatewaySettings#binding_status}
	BindingStatus *string `field:"optional" json:"bindingStatus" yaml:"bindingStatus"`
	// Specify the UUID of the certificate (ID from MTLS certificate store).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#id ZeroTrustGatewaySettings#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#updated_at ZeroTrustGatewaySettings#updated_at}.
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

