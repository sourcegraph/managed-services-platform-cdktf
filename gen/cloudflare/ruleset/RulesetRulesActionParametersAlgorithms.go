package ruleset


type RulesetRulesActionParametersAlgorithms struct {
	// Name of the compression algorithm to use. Available values: `zstd`, `gzip`, `brotli`, `auto`, `default`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

