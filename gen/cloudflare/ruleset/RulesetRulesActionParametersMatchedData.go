package ruleset


type RulesetRulesActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#public_key Ruleset#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

