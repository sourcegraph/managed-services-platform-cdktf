package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAdd struct {
	// The name of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_compute_region_url_map#header_name GoogleComputeRegionUrlMap#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_compute_region_url_map#header_value GoogleComputeRegionUrlMap#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
	// If false, headerValue is appended to any values that already exist for the header.
	//
	// If true, headerValue is set for the header, discarding any values that were set for that header.
	// The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_compute_region_url_map#replace GoogleComputeRegionUrlMap#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
}

