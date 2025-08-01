package computerouterroutepolicy


type ComputeRouterRoutePolicyTerms struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_router_route_policy#match ComputeRouterRoutePolicy#match}
	Match *ComputeRouterRoutePolicyTermsMatch `field:"required" json:"match" yaml:"match"`
	// The evaluation priority for this term, which must be between 0 (inclusive) and 231 (exclusive), and unique within the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_router_route_policy#priority ComputeRouterRoutePolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_router_route_policy#actions ComputeRouterRoutePolicy#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
}

