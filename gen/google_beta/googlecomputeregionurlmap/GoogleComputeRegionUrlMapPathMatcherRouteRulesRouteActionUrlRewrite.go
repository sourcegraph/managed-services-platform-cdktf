package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite struct {
	// Prior to forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite.
	//
	// The value must be between 1 and
	// 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#host_rewrite GoogleComputeRegionUrlMap#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Prior to forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite.
	//
	// The value must
	// be between 1 and 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#path_prefix_rewrite GoogleComputeRegionUrlMap#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
	// Prior to forwarding the request to the selected origin, if the request matched a pathTemplateMatch, the matching portion of the request's path is replaced re-written using the pattern specified by pathTemplateRewrite.
	//
	// pathTemplateRewrite must be between 1 and 255 characters
	// (inclusive), must start with a '/', and must only use variables
	// captured by the route's pathTemplate matchers.
	//
	// pathTemplateRewrite may only be used when all of a route's
	// MatchRules specify pathTemplate.
	//
	// Only one of pathPrefixRewrite and pathTemplateRewrite may be
	// specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#path_template_rewrite GoogleComputeRegionUrlMap#path_template_rewrite}
	PathTemplateRewrite *string `field:"optional" json:"pathTemplateRewrite" yaml:"pathTemplateRewrite"`
}

