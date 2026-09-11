package datacloudflarecustomorigintruststore


type DataCloudflareCustomOriginTrustStoreFilter struct {
	// Limit to the number of records returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_store#limit DataCloudflareCustomOriginTrustStore#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_store#offset DataCloudflareCustomOriginTrustStore#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

