package computeregionurlmap


type ComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite struct {
	// Prior to forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite.
	//
	// The value must be between 1 and
	// 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_url_map#host_rewrite ComputeRegionUrlMap#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Prior to forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite.
	//
	// The value must
	// be between 1 and 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_url_map#path_prefix_rewrite ComputeRegionUrlMap#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
}
