package datacloudflarezerotrustdextest


type DataCloudflareZeroTrustDexTestFilter struct {
	// Filter by test type Available values: "http", "traceroute".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/zero_trust_dex_test#kind DataCloudflareZeroTrustDexTest#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Filter by test name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/zero_trust_dex_test#test_name DataCloudflareZeroTrustDexTest#test_name}
	TestName *string `field:"optional" json:"testName" yaml:"testName"`
}

