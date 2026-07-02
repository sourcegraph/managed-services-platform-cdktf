package ruleset

type RulesetRulesActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#device_type Ruleset#device_type}
	DeviceType any `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Use the user agents's country in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#geo Ruleset#geo}
	Geo any `field:"optional" json:"geo" yaml:"geo"`
	// Use the user agent's language in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#lang Ruleset#lang}
	Lang any `field:"optional" json:"lang" yaml:"lang"`
}
