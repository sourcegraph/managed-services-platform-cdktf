package ruleset


type RulesetRulesActionParametersSni struct {
	// The SNI override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

