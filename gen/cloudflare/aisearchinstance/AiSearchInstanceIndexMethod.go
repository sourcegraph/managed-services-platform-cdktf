package aisearchinstance


type AiSearchInstanceIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#keyword AiSearchInstance#keyword}
	Keyword interface{} `field:"required" json:"keyword" yaml:"keyword"`
	// Enable vector (embedding) storage backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#vector AiSearchInstance#vector}
	Vector interface{} `field:"required" json:"vector" yaml:"vector"`
}

