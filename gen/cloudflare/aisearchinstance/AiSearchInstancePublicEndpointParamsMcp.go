package aisearchinstance


type AiSearchInstancePublicEndpointParamsMcp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#description AiSearchInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disable MCP endpoint for this public endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_search_instance#disabled AiSearchInstance#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

