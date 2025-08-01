package datacloudflarezerotrustlist


type DataCloudflareZeroTrustListFilter struct {
	// The type of list. Available values: "SERIAL", "URL", "DOMAIN", "EMAIL", "IP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_list#type DataCloudflareZeroTrustList#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

