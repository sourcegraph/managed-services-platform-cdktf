package networkservicestcproute


type NetworkServicesTcpRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tcp_route#action NetworkServicesTcpRoute#action}
	Action *NetworkServicesTcpRouteRulesAction `field:"required" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tcp_route#matches NetworkServicesTcpRoute#matches}
	Matches interface{} `field:"optional" json:"matches" yaml:"matches"`
}

