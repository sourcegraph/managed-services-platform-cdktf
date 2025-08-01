package emailroutingcatchall


type EmailRoutingCatchAllActions struct {
	// Type of action for catch-all rule. Available values: "drop", "forward", "worker".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/email_routing_catch_all#value EmailRoutingCatchAll#value}.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

