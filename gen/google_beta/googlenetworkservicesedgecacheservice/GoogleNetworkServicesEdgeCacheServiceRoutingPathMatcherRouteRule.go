package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRule struct {
	// match_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#match_rule GoogleNetworkServicesEdgeCacheService#match_rule}
	MatchRule interface{} `field:"required" json:"matchRule" yaml:"matchRule"`
	// The priority of this route rule, where 1 is the highest priority.
	//
	// You cannot configure two or more routeRules with the same priority. Priority for each rule must be set to a number between 1 and 999 inclusive.
	//
	// Priority numbers can have gaps, which enable you to add or remove rules in the future without affecting the rest of the rules. For example, 1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers
	// to which you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the future without any impact on existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#priority GoogleNetworkServicesEdgeCacheService#priority}
	Priority *string `field:"required" json:"priority" yaml:"priority"`
	// A human-readable description of the routeRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#description GoogleNetworkServicesEdgeCacheService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#header_action GoogleNetworkServicesEdgeCacheService#header_action}
	HeaderAction *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// The Origin resource that requests to this route should fetch from when a matching response is not in cache.
	//
	// Origins can be defined as short names ("my-origin") or fully-qualified resource URLs - e.g. "networkservices.googleapis.com/projects/my-project/global/edgecacheorigins/my-origin"
	//
	// Only one of origin or urlRedirect can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#origin GoogleNetworkServicesEdgeCacheService#origin}
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#route_action GoogleNetworkServicesEdgeCacheService#route_action}
	RouteAction *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction `field:"optional" json:"routeAction" yaml:"routeAction"`
	// route_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#route_methods GoogleNetworkServicesEdgeCacheService#route_methods}
	RouteMethods *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteMethods `field:"optional" json:"routeMethods" yaml:"routeMethods"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#url_redirect GoogleNetworkServicesEdgeCacheService#url_redirect}
	UrlRedirect *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
}

