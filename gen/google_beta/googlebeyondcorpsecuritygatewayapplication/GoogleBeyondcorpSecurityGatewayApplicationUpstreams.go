package googlebeyondcorpsecuritygatewayapplication


type GoogleBeyondcorpSecurityGatewayApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_beyondcorp_security_gateway_application#egress_policy GoogleBeyondcorpSecurityGatewayApplication#egress_policy}
	EgressPolicy *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_beyondcorp_security_gateway_application#network GoogleBeyondcorpSecurityGatewayApplication#network}
	Network *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
}

