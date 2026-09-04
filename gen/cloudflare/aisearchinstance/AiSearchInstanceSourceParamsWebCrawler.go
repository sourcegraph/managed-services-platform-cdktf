package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#parse_options AiSearchInstance#parse_options}.
	ParseOptions *AiSearchInstanceSourceParamsWebCrawlerParseOptions `field:"optional" json:"parseOptions" yaml:"parseOptions"`
	// Available values: "sitemap", "feed-rss".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#parse_type AiSearchInstance#parse_type}
	ParseType *string `field:"optional" json:"parseType" yaml:"parseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#store_options AiSearchInstance#store_options}.
	StoreOptions *AiSearchInstanceSourceParamsWebCrawlerStoreOptions `field:"optional" json:"storeOptions" yaml:"storeOptions"`
}

