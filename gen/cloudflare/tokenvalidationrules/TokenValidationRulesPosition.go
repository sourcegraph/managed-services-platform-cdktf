package tokenvalidationrules


type TokenValidationRulesPosition struct {
	// Move rule to after rule with this ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#after TokenValidationRules#after}
	After *string `field:"optional" json:"after" yaml:"after"`
	// Move rule to before rule with this ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#before TokenValidationRules#before}
	Before *string `field:"optional" json:"before" yaml:"before"`
	// Move rule to this position.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#index TokenValidationRules#index}
	Index *float64 `field:"optional" json:"index" yaml:"index"`
}

