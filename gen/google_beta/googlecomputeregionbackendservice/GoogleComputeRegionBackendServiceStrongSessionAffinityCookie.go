package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceStrongSessionAffinityCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_region_backend_service#name GoogleComputeRegionBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_region_backend_service#path GoogleComputeRegionBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_region_backend_service#ttl GoogleComputeRegionBackendService#ttl}
	Ttl *GoogleComputeRegionBackendServiceStrongSessionAffinityCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

