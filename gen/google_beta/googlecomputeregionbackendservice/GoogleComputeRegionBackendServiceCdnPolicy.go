package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceCdnPolicy struct {
	// cache_key_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#cache_key_policy GoogleComputeRegionBackendService#cache_key_policy}
	CacheKeyPolicy *GoogleComputeRegionBackendServiceCdnPolicyCacheKeyPolicy `field:"optional" json:"cacheKeyPolicy" yaml:"cacheKeyPolicy"`
	// Specifies the cache setting for all responses from this backend.
	//
	// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#cache_mode GoogleComputeRegionBackendService#cache_mode}
	CacheMode *string `field:"optional" json:"cacheMode" yaml:"cacheMode"`
	// Specifies the maximum allowed TTL for cached content served by this origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#client_ttl GoogleComputeRegionBackendService#client_ttl}
	ClientTtl *float64 `field:"optional" json:"clientTtl" yaml:"clientTtl"`
	// Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL (max-age or s-max-age).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#default_ttl GoogleComputeRegionBackendService#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Specifies the maximum allowed TTL for cached content served by this origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#max_ttl GoogleComputeRegionBackendService#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#negative_caching GoogleComputeRegionBackendService#negative_caching}
	NegativeCaching interface{} `field:"optional" json:"negativeCaching" yaml:"negativeCaching"`
	// negative_caching_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#negative_caching_policy GoogleComputeRegionBackendService#negative_caching_policy}
	NegativeCachingPolicy interface{} `field:"optional" json:"negativeCachingPolicy" yaml:"negativeCachingPolicy"`
	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#serve_while_stale GoogleComputeRegionBackendService#serve_while_stale}
	ServeWhileStale *float64 `field:"optional" json:"serveWhileStale" yaml:"serveWhileStale"`
	// Maximum number of seconds the response to a signed URL request will be considered fresh, defaults to 1hr (3600s).
	//
	// After this
	// time period, the response will be revalidated before
	// being served.
	//
	// When serving responses to signed URL requests, Cloud CDN will
	// internally behave as though all responses from this backend had a
	// "Cache-Control: public, max-age=[TTL]" header, regardless of any
	// existing Cache-Control header. The actual headers served in
	// responses will not be altered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_backend_service#signed_url_cache_max_age_sec GoogleComputeRegionBackendService#signed_url_cache_max_age_sec}
	SignedUrlCacheMaxAgeSec *float64 `field:"optional" json:"signedUrlCacheMaxAgeSec" yaml:"signedUrlCacheMaxAgeSec"`
}

