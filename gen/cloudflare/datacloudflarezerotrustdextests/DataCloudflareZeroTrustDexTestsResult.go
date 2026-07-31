package datacloudflarezerotrustdextests


type DataCloudflareZeroTrustDexTestsResult struct {
	// DEX rules targeted by this test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/zero_trust_dex_tests#target_policies DataCloudflareZeroTrustDexTests#target_policies}
	TargetPolicies interface{} `field:"optional" json:"targetPolicies" yaml:"targetPolicies"`
}

