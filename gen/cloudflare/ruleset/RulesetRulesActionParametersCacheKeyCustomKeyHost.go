package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHost struct {
	// Whether to use the resolved host in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#resolved Ruleset#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}

