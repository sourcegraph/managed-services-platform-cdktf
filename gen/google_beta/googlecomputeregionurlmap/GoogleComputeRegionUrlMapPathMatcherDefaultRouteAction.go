package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherDefaultRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#cors_policy GoogleComputeRegionUrlMap#cors_policy}
	CorsPolicy *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#fault_injection_policy GoogleComputeRegionUrlMap#fault_injection_policy}
	FaultInjectionPolicy *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// max_stream_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#max_stream_duration GoogleComputeRegionUrlMap#max_stream_duration}
	MaxStreamDuration *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration `field:"optional" json:"maxStreamDuration" yaml:"maxStreamDuration"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#request_mirror_policy GoogleComputeRegionUrlMap#request_mirror_policy}
	RequestMirrorPolicy *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#retry_policy GoogleComputeRegionUrlMap#retry_policy}
	RetryPolicy *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#timeout GoogleComputeRegionUrlMap#timeout}
	Timeout *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#url_rewrite GoogleComputeRegionUrlMap#url_rewrite}
	UrlRewrite *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#weighted_backend_services GoogleComputeRegionUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

