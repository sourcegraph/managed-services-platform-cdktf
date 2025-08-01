package ruleset


type RulesetRulesActionParametersUriQuery struct {
	// Expression to evaluate for the replacement value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Predefined replacement value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

