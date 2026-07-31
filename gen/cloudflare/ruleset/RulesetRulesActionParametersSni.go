package ruleset


type RulesetRulesActionParametersSni struct {
	// A value to override the SNI to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

