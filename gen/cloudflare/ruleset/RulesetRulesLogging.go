package ruleset


type RulesetRulesLogging struct {
	// Whether to generate a log when the rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

