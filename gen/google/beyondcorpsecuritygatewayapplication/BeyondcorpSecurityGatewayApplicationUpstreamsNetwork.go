package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreamsNetwork struct {
	// Required. Network name is of the format: 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application#name BeyondcorpSecurityGatewayApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

