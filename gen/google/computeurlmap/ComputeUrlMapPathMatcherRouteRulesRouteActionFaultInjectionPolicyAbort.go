package computeurlmap


type ComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request. The value must be between 200 and 599 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_url_map#http_status ComputeUrlMap#http_status}
	HttpStatus *float64 `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection.
	//
	// The value must be between 0.0 and 100.0
	// inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_url_map#percentage ComputeUrlMap#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

