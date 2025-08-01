package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key.
	//
	// A value of true will use the resolved host, while a value or false will use the original host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#resolved Ruleset#resolved}
	Resolved interface{} `field:"optional" json:"resolved" yaml:"resolved"`
}

