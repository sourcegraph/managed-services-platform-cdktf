package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryString struct {
	// A list of query string parameters NOT used to build the cache key.
	//
	// All parameters present in the request but missing in this list will be used to build the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#exclude Ruleset#exclude}
	Exclude *RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// A list of query string parameters used to build the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#include Ruleset#include}
	Include *RulesetRulesActionParametersCacheKeyCustomKeyQueryStringInclude `field:"optional" json:"include" yaml:"include"`
}

