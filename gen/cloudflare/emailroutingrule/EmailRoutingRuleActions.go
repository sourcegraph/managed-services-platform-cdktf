package emailroutingrule


type EmailRoutingRuleActions struct {
	// Type of supported action. Available values: "drop", "forward", "worker".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_routing_rule#value EmailRoutingRule#value}.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

