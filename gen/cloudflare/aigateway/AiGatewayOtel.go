package aigateway


type AiGatewayOtel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway#authorization AiGateway#authorization}.
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway#headers AiGateway#headers}.
	Headers *map[string]*string `field:"required" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway#url AiGateway#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Available values: "json", "protobuf".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway#content_type AiGateway#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

