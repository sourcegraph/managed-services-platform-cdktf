package aigatewaydynamicrouting


type AiGatewayDynamicRoutingElementsOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#element_id AiGatewayDynamicRouting#element_id}.
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#fallback AiGatewayDynamicRouting#fallback}.
	Fallback *AiGatewayDynamicRoutingElementsOutputsFallback `field:"optional" json:"fallback" yaml:"fallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#false AiGatewayDynamicRouting#false}.
	False *AiGatewayDynamicRoutingElementsOutputsFalse `field:"optional" json:"false" yaml:"false"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#next AiGatewayDynamicRouting#next}.
	Next *AiGatewayDynamicRoutingElementsOutputsNext `field:"optional" json:"next" yaml:"next"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#success AiGatewayDynamicRouting#success}.
	Success *AiGatewayDynamicRoutingElementsOutputsSuccess `field:"optional" json:"success" yaml:"success"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ai_gateway_dynamic_routing#true AiGatewayDynamicRouting#true}.
	True *AiGatewayDynamicRoutingElementsOutputsTrue `field:"optional" json:"true" yaml:"true"`
}

