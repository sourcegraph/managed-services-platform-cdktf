package networkserviceshttproute


type NetworkServicesHttpRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#action NetworkServicesHttpRoute#action}
	Action *NetworkServicesHttpRouteRulesAction `field:"optional" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_services_http_route#matches NetworkServicesHttpRoute#matches}
	Matches interface{} `field:"optional" json:"matches" yaml:"matches"`
}

