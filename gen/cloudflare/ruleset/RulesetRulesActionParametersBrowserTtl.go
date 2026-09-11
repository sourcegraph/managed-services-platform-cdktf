package ruleset


type RulesetRulesActionParametersBrowserTtl struct {
	// The browser TTL mode. Available values: "respect_origin", "bypass_by_default", "override_origin", "bypass".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The browser TTL (in seconds) if you choose the "override_origin" mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
}

