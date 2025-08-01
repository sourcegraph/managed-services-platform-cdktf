package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRules struct {
	// For routeRules within a given pathMatcher, priority determines the order in which load balancer will interpret routeRules.
	//
	// RouteRules are evaluated
	// in order of priority, from the lowest to highest number. The priority of
	// a rule decreases as its number increases (1, 2, 3, N+1). The first rule
	// that matches the request is applied.
	//
	// You cannot configure two or more routeRules with the same priority.
	// Priority for each rule must be set to a number between 0 and
	// 2147483647 inclusive.
	//
	// Priority numbers can have gaps, which enable you to add or remove rules
	// in the future without affecting the rest of the rules. For example,
	// 1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers to which
	// you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the
	// future without any impact on existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#priority GoogleComputeRegionUrlMap#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#header_action GoogleComputeRegionUrlMap#header_action}
	HeaderAction *GoogleComputeRegionUrlMapPathMatcherRouteRulesHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// match_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#match_rules GoogleComputeRegionUrlMap#match_rules}
	MatchRules interface{} `field:"optional" json:"matchRules" yaml:"matchRules"`
	// route_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#route_action GoogleComputeRegionUrlMap#route_action}
	RouteAction *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction `field:"optional" json:"routeAction" yaml:"routeAction"`
	// The region backend service resource to which traffic is directed if this rule is matched.
	//
	// If routeAction is additionally specified,
	// advanced routing actions like URL Rewrites, etc. take effect prior to sending
	// the request to the backend. However, if service is specified, routeAction cannot
	// contain any weightedBackendService s. Conversely, if routeAction specifies any
	// weightedBackendServices, service must not be specified. Only one of urlRedirect,
	// service or routeAction.weightedBackendService must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#service GoogleComputeRegionUrlMap#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#url_redirect GoogleComputeRegionUrlMap#url_redirect}
	UrlRedirect *GoogleComputeRegionUrlMapPathMatcherRouteRulesUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
}

