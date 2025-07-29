package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationEndpointMatchers struct {
	// Required. Hostname of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_beyondcorp_security_gateway_application#hostname GoogleBeyondcorpSecurityGatewayApplication#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Optional. Ports of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_beyondcorp_security_gateway_application#ports GoogleBeyondcorpSecurityGatewayApplication#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

