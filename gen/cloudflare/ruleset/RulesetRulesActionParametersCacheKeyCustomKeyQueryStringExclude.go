package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Whether to exclude all query string parameters from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#all Ruleset#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// A list of query string parameters to exclude from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#list Ruleset#list}
	List *[]*string `field:"optional" json:"list" yaml:"list"`
}

