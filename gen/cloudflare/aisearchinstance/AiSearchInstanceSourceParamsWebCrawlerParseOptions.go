package aisearchinstance


type AiSearchInstanceSourceParamsWebCrawlerParseOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#include_headers AiSearchInstance#include_headers}.
	IncludeHeaders *map[string]*string `field:"optional" json:"includeHeaders" yaml:"includeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#include_images AiSearchInstance#include_images}.
	IncludeImages interface{} `field:"optional" json:"includeImages" yaml:"includeImages"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is 'sitemap'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#specific_sitemaps AiSearchInstance#specific_sitemaps}
	SpecificSitemaps *[]*string `field:"optional" json:"specificSitemaps" yaml:"specificSitemaps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#use_browser_rendering AiSearchInstance#use_browser_rendering}.
	UseBrowserRendering interface{} `field:"optional" json:"useBrowserRendering" yaml:"useBrowserRendering"`
}

