package networkservicestlsroute


type NetworkServicesTlsRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tls_route#action NetworkServicesTlsRoute#action}
	Action *NetworkServicesTlsRouteRulesAction `field:"required" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_tls_route#matches NetworkServicesTlsRoute#matches}
	Matches interface{} `field:"required" json:"matches" yaml:"matches"`
}

