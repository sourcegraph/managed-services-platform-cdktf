package ruleset


type RulesetRulesActionParametersMustUnderstand struct {
	// The operation to perform. Available values: "set", "remove".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#operation Ruleset#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Whether to apply the directive only to Cloudflare's cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/ruleset#cloudflare_only Ruleset#cloudflare_only}
	CloudflareOnly interface{} `field:"optional" json:"cloudflareOnly" yaml:"cloudflareOnly"`
}

