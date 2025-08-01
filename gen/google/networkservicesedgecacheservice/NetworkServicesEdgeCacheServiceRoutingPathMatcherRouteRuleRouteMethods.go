package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteMethods struct {
	// The non-empty set of HTTP methods that are allowed for this route.
	//
	// Any combination of "GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", and "PATCH".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_edge_cache_service#allowed_methods NetworkServicesEdgeCacheService#allowed_methods}
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
}

