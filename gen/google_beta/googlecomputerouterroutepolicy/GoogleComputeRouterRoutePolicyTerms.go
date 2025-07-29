package googlecomputerouterroutepolicy


type GoogleComputeRouterRoutePolicyTerms struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_route_policy#match GoogleComputeRouterRoutePolicy#match}
	Match *GoogleComputeRouterRoutePolicyTermsMatch `field:"required" json:"match" yaml:"match"`
	// The evaluation priority for this term, which must be between 0 (inclusive) and 231 (exclusive), and unique within the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_route_policy#priority GoogleComputeRouterRoutePolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router_route_policy#actions GoogleComputeRouterRoutePolicy#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
}

