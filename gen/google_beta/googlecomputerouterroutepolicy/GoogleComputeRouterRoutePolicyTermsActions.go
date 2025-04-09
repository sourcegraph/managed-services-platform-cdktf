package googlecomputerouterroutepolicy


type GoogleComputeRouterRoutePolicyTermsActions struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_router_route_policy#expression GoogleComputeRouterRoutePolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_router_route_policy#description GoogleComputeRouterRoutePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_router_route_policy#location GoogleComputeRouterRoutePolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_router_route_policy#title GoogleComputeRouterRoutePolicy#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

