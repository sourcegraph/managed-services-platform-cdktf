package ruleset

type RulesetRulesActionParametersAutominify struct {
	// Minify CSS files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#css Ruleset#css}
	Css any `field:"optional" json:"css" yaml:"css"`
	// Minify HTML files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#html Ruleset#html}
	Html any `field:"optional" json:"html" yaml:"html"`
	// Minify JS files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#js Ruleset#js}
	Js any `field:"optional" json:"js" yaml:"js"`
}
