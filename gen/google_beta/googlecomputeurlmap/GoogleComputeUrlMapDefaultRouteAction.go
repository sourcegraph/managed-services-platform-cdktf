package googlecomputeurlmap


type GoogleComputeUrlMapDefaultRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#cors_policy GoogleComputeUrlMap#cors_policy}
	CorsPolicy *GoogleComputeUrlMapDefaultRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#fault_injection_policy GoogleComputeUrlMap#fault_injection_policy}
	FaultInjectionPolicy *GoogleComputeUrlMapDefaultRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// max_stream_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#max_stream_duration GoogleComputeUrlMap#max_stream_duration}
	MaxStreamDuration *GoogleComputeUrlMapDefaultRouteActionMaxStreamDuration `field:"optional" json:"maxStreamDuration" yaml:"maxStreamDuration"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#request_mirror_policy GoogleComputeUrlMap#request_mirror_policy}
	RequestMirrorPolicy *GoogleComputeUrlMapDefaultRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#retry_policy GoogleComputeUrlMap#retry_policy}
	RetryPolicy *GoogleComputeUrlMapDefaultRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#timeout GoogleComputeUrlMap#timeout}
	Timeout *GoogleComputeUrlMapDefaultRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#url_rewrite GoogleComputeUrlMap#url_rewrite}
	UrlRewrite *GoogleComputeUrlMapDefaultRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#weighted_backend_services GoogleComputeUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

