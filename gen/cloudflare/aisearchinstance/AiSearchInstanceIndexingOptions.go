package aisearchinstance


type AiSearchInstanceIndexingOptions struct {
	// Tokenizer used for keyword search indexing.
	//
	// porter provides word-level tokenization with Porter stemming (good for natural language queries). trigram enables character-level substring matching (good for partial matches, code, identifiers). Changing this triggers a full re-index. Defaults to porter.
	// Available values: "porter", "trigram".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#keyword_tokenizer AiSearchInstance#keyword_tokenizer}
	KeywordTokenizer *string `field:"optional" json:"keywordTokenizer" yaml:"keywordTokenizer"`
}

