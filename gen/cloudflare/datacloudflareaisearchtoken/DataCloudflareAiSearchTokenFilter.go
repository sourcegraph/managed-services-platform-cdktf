package datacloudflareaisearchtoken


type DataCloudflareAiSearchTokenFilter struct {
	// Filter tokens whose name contains this string (case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_token#search DataCloudflareAiSearchToken#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

