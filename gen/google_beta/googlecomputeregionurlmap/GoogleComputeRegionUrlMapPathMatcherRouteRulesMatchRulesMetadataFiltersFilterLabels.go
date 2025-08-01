package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabels struct {
	// Name of metadata label.
	//
	// The name can have a maximum length of 1024 characters
	// and must be at least 1 character long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#name GoogleComputeRegionUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the label must match the specified value. value can have a maximum length of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#value GoogleComputeRegionUrlMap#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

