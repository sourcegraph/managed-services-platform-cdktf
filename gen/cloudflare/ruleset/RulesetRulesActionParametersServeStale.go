package ruleset


type RulesetRulesActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating.
	//
	// If true, Cloudflare will not serve stale content while getting the latest content from the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/ruleset#disable_stale_while_updating Ruleset#disable_stale_while_updating}
	DisableStaleWhileUpdating interface{} `field:"required" json:"disableStaleWhileUpdating" yaml:"disableStaleWhileUpdating"`
}

