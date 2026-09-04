package ruleset


type RulesetRulesActionParametersAlgorithms struct {
	// Name of the compression algorithm to enable. Available values: "none", "auto", "default", "gzip", "brotli", "zstd".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

