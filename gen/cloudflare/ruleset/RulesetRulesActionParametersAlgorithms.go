package ruleset


type RulesetRulesActionParametersAlgorithms struct {
	// Name of compression algorithm to enable. Available values: "none", "auto", "default", "gzip", "brotli", "zstd".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

