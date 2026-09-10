package aisearchinstance


type AiSearchInstanceRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by.
	//
	// Use 'timestamp' for document freshness, or any custom_metadata field. Numeric and datetime fields support asc/desc directions; text/boolean fields support exists/not_exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#field AiSearchInstance#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Boost direction.
	//
	// 'desc' = higher values rank higher (e.g. newer timestamps). 'asc' = lower values rank higher. 'exists' = boost chunks that have the field. 'not_exists' = boost chunks that lack the field. Optional - defaults to 'asc' for numeric/datetime fields, 'exists' for text/boolean fields.
	// Available values: "asc", "desc", "exists", "not_exists".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#direction AiSearchInstance#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

