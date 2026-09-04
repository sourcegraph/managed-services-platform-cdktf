package emailroutingcatchall


type EmailRoutingCatchAllMatchers struct {
	// Type of matcher. Default is 'all'. Available values: "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

