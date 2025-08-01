package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceConsistentHash struct {
	// http_cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#http_cookie GoogleComputeRegionBackendService#http_cookie}
	HttpCookie *GoogleComputeRegionBackendServiceConsistentHashHttpCookie `field:"optional" json:"httpCookie" yaml:"httpCookie"`
	// The hash based on the value of the specified header field.
	//
	// This field is applicable if the sessionAffinity is set to HEADER_FIELD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#http_header_name GoogleComputeRegionBackendService#http_header_name}
	HttpHeaderName *string `field:"optional" json:"httpHeaderName" yaml:"httpHeaderName"`
	// The minimum number of virtual nodes to use for the hash ring.
	//
	// Larger ring sizes result in more granular load
	// distributions. If the number of hosts in the load balancing pool
	// is larger than the ring size, each host will be assigned a single
	// virtual node.
	// Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#minimum_ring_size GoogleComputeRegionBackendService#minimum_ring_size}
	MinimumRingSize *float64 `field:"optional" json:"minimumRingSize" yaml:"minimumRingSize"`
}

