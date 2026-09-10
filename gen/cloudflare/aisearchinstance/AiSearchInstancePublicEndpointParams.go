package aisearchinstance


type AiSearchInstancePublicEndpointParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#authorized_hosts AiSearchInstance#authorized_hosts}.
	AuthorizedHosts *[]*string `field:"optional" json:"authorizedHosts" yaml:"authorizedHosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#chat_completions_endpoint AiSearchInstance#chat_completions_endpoint}.
	ChatCompletionsEndpoint *AiSearchInstancePublicEndpointParamsChatCompletionsEndpoint `field:"optional" json:"chatCompletionsEndpoint" yaml:"chatCompletionsEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#enabled AiSearchInstance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#mcp AiSearchInstance#mcp}.
	Mcp *AiSearchInstancePublicEndpointParamsMcp `field:"optional" json:"mcp" yaml:"mcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#rate_limit AiSearchInstance#rate_limit}.
	RateLimit *AiSearchInstancePublicEndpointParamsRateLimit `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#search_endpoint AiSearchInstance#search_endpoint}.
	SearchEndpoint *AiSearchInstancePublicEndpointParamsSearchEndpoint `field:"optional" json:"searchEndpoint" yaml:"searchEndpoint"`
}

