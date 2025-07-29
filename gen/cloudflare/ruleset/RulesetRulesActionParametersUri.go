package ruleset


type RulesetRulesActionParametersUri struct {
	// Path portion rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#path Ruleset#path}
	Path *RulesetRulesActionParametersUriPath `field:"optional" json:"path" yaml:"path"`
	// Query portion rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#query Ruleset#query}
	Query *RulesetRulesActionParametersUriQuery `field:"optional" json:"query" yaml:"query"`
}

