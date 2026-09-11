package aisearchinstance


type AiSearchInstanceRetrievalOptions struct {
	// Metadata fields to boost search results by.
	//
	// Each entry specifies a metadata field and an optional direction. Direction defaults to 'asc' for numeric fields and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined custom_metadata field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#boost_by AiSearchInstance#boost_by}
	BoostBy interface{} `field:"optional" json:"boostBy" yaml:"boostBy"`
	// Controls which documents are candidates for BM25 scoring.
	//
	// 'and' restricts candidates to documents containing all query terms; 'or' includes any document containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
	// Available values: "and", "or".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#keyword_match_mode AiSearchInstance#keyword_match_mode}
	KeywordMatchMode *string `field:"optional" json:"keywordMatchMode" yaml:"keywordMatchMode"`
}

