package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_url_map#header_name GoogleComputeUrlMap#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_url_map#header_value GoogleComputeUrlMap#header_value}
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
	// If false, headerValue is appended to any values that already exist for the header.
	//
	// If true, headerValue is set for the header, discarding any values that
	// were set for that header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_url_map#replace GoogleComputeUrlMap#replace}
	Replace interface{} `field:"required" json:"replace" yaml:"replace"`
}

