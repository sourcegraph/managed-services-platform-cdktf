package networkservicestcproute


type NetworkServicesTcpRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tcp_route#create NetworkServicesTcpRoute#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tcp_route#delete NetworkServicesTcpRoute#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tcp_route#update NetworkServicesTcpRoute#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

