package datacloudflareaisearchinstance


type DataCloudflareAiSearchInstanceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance#namespace DataCloudflareAiSearchInstance#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Order By Column Name Available values: "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance#order_by DataCloudflareAiSearchInstance#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Order By Direction Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance#order_by_direction DataCloudflareAiSearchInstance#order_by_direction}
	OrderByDirection *string `field:"optional" json:"orderByDirection" yaml:"orderByDirection"`
	// Search by id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance#search DataCloudflareAiSearchInstance#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

