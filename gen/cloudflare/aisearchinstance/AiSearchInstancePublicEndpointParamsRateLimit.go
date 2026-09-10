package aisearchinstance


type AiSearchInstancePublicEndpointParamsRateLimit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#period_ms AiSearchInstance#period_ms}.
	PeriodMs *float64 `field:"optional" json:"periodMs" yaml:"periodMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#requests AiSearchInstance#requests}.
	Requests *float64 `field:"optional" json:"requests" yaml:"requests"`
	// Available values: "fixed", "sliding".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#technique AiSearchInstance#technique}
	Technique *string `field:"optional" json:"technique" yaml:"technique"`
}

