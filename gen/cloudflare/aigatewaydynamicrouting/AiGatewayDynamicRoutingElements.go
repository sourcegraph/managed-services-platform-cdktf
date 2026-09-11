package aigatewaydynamicrouting


type AiGatewayDynamicRoutingElements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#id AiGatewayDynamicRouting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#outputs AiGatewayDynamicRouting#outputs}.
	Outputs *AiGatewayDynamicRoutingElementsOutputs `field:"required" json:"outputs" yaml:"outputs"`
	// Available values: "start", "conditional", "percentage", "rate", "model", "end".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#type AiGatewayDynamicRouting#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#properties AiGatewayDynamicRouting#properties}.
	Properties *AiGatewayDynamicRoutingElementsProperties `field:"optional" json:"properties" yaml:"properties"`
}

