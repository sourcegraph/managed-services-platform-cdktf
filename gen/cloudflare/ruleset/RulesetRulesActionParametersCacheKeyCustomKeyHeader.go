package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHeader struct {
	// List of headers to check for presence in the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/ruleset#check_presence Ruleset#check_presence}
	CheckPresence *[]*string `field:"optional" json:"checkPresence" yaml:"checkPresence"`
	// Dictionary of headers mapping to lists of values to check for presence in the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/ruleset#contains Ruleset#contains}
	Contains interface{} `field:"optional" json:"contains" yaml:"contains"`
	// Exclude the origin header from the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/ruleset#exclude_origin Ruleset#exclude_origin}
	ExcludeOrigin interface{} `field:"optional" json:"excludeOrigin" yaml:"excludeOrigin"`
	// List of headers to include in the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/ruleset#include Ruleset#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

