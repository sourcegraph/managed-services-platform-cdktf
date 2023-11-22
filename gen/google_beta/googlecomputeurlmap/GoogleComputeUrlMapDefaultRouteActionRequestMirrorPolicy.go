package googlecomputeurlmap


type GoogleComputeUrlMapDefaultRouteActionRequestMirrorPolicy struct {
	// The full or partial URL to the BackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_compute_url_map#backend_service GoogleComputeUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

