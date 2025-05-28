package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyDelay struct {
	// Specify a fixed delay before forwarding the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_network_services_http_route#fixed_delay GoogleNetworkServicesHttpRoute#fixed_delay}
	FixedDelay *string `field:"optional" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic on which delay will be injected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_network_services_http_route#percentage GoogleNetworkServicesHttpRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

