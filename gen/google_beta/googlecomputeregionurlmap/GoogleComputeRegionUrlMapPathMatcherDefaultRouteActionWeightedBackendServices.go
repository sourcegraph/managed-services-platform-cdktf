package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServices struct {
	// The full or partial URL to the default BackendService resource.
	//
	// Before forwarding the
	// request to backendService, the loadbalancer applies any relevant headerActions
	// specified as part of this backendServiceWeight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#backend_service GoogleComputeRegionUrlMap#backend_service}
	BackendService *string `field:"optional" json:"backendService" yaml:"backendService"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#header_action GoogleComputeRegionUrlMap#header_action}
	HeaderAction *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// Specifies the fraction of traffic sent to backendService, computed as weight / (sum of all weightedBackendService weights in routeAction) .
	//
	// The selection of a backend service is determined only for new traffic. Once a user's request
	// has been directed to a backendService, subsequent requests will be sent to the same backendService
	// as determined by the BackendService's session affinity policy.
	//
	// The value must be between 0 and 1000
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_url_map#weight GoogleComputeRegionUrlMap#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

