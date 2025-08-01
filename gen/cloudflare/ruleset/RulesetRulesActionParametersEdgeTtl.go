package ruleset


type RulesetRulesActionParametersEdgeTtl struct {
	// Edge TTL options. Available values: "respect_origin", "bypass_by_default", "override_origin".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The TTL (in seconds) if you choose override_origin mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// List of single status codes, or status code ranges to apply the selected mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#status_code_ttl Ruleset#status_code_ttl}
	StatusCodeTtl interface{} `field:"optional" json:"statusCodeTtl" yaml:"statusCodeTtl"`
}

