package googlenetworkservicesedgecacheservice

type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction struct {
	// request_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#request_header_to_add GoogleNetworkServicesEdgeCacheService#request_header_to_add}
	RequestHeaderToAdd any `field:"optional" json:"requestHeaderToAdd" yaml:"requestHeaderToAdd"`
	// request_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#request_header_to_remove GoogleNetworkServicesEdgeCacheService#request_header_to_remove}
	RequestHeaderToRemove any `field:"optional" json:"requestHeaderToRemove" yaml:"requestHeaderToRemove"`
	// response_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#response_header_to_add GoogleNetworkServicesEdgeCacheService#response_header_to_add}
	ResponseHeaderToAdd any `field:"optional" json:"responseHeaderToAdd" yaml:"responseHeaderToAdd"`
	// response_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#response_header_to_remove GoogleNetworkServicesEdgeCacheService#response_header_to_remove}
	ResponseHeaderToRemove any `field:"optional" json:"responseHeaderToRemove" yaml:"responseHeaderToRemove"`
}
