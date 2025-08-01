package ruleset


type RulesetRulesActionParametersFromValueTargetUrl struct {
	// An expression to evaluate to get the URL to redirect the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The URL to redirect the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

