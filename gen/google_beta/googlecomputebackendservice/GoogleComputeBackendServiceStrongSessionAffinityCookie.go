package googlecomputebackendservice


type GoogleComputeBackendServiceStrongSessionAffinityCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_backend_service#name GoogleComputeBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_backend_service#path GoogleComputeBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_backend_service#ttl GoogleComputeBackendService#ttl}
	Ttl *GoogleComputeBackendServiceStrongSessionAffinityCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

