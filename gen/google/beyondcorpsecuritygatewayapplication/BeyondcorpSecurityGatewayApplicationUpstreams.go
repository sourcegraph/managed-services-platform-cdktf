package beyondcorpsecuritygatewayapplication


type BeyondcorpSecurityGatewayApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application#egress_policy BeyondcorpSecurityGatewayApplication#egress_policy}
	EgressPolicy *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application#network BeyondcorpSecurityGatewayApplication#network}
	Network *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
}

