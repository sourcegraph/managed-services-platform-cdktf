package aisearchinstance


type AiSearchInstanceSourceParams struct {
	// List of path patterns to exclude.
	//
	// Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /admin/** matches /admin/users and /admin/settings/advanced)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#exclude_items AiSearchInstance#exclude_items}
	ExcludeItems *[]*string `field:"optional" json:"excludeItems" yaml:"excludeItems"`
	// List of path patterns to include.
	//
	// Uses micromatch glob syntax: * matches within a path segment, ** matches across path segments (e.g., /blog/** matches /blog/post and /blog/2024/post)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#include_items AiSearchInstance#include_items}
	IncludeItems *[]*string `field:"optional" json:"includeItems" yaml:"includeItems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#prefix AiSearchInstance#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#r2_jurisdiction AiSearchInstance#r2_jurisdiction}.
	R2Jurisdiction *string `field:"optional" json:"r2Jurisdiction" yaml:"r2Jurisdiction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#web_crawler AiSearchInstance#web_crawler}.
	WebCrawler *AiSearchInstanceSourceParamsWebCrawler `field:"optional" json:"webCrawler" yaml:"webCrawler"`
}

