package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawlerCrawlOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#depth AiSearchInstance#depth}.
	Depth *float64 `field:"optional" json:"depth" yaml:"depth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#include_external_links AiSearchInstance#include_external_links}.
	IncludeExternalLinks interface{} `field:"optional" json:"includeExternalLinks" yaml:"includeExternalLinks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#include_subdomains AiSearchInstance#include_subdomains}.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#max_age AiSearchInstance#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Available values: "all", "sitemaps", "links".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#source AiSearchInstance#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

