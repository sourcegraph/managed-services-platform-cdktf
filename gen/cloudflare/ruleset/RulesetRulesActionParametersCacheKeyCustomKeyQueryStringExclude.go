package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Determines whether to exclude all query string parameters from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#all Ruleset#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#list Ruleset#list}.
	List *[]*string `field:"optional" json:"list" yaml:"list"`
}

