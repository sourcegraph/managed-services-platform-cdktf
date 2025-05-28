package teamsrule


type TeamsRuleRuleSettingsResolveDnsInternally struct {
	// The fallback behavior to apply when the internal DNS response code is different from 'NOERROR' or when the response data only contains CNAME records for 'A' or 'AAAA' queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_rule#fallback TeamsRule#fallback}
	Fallback *string `field:"optional" json:"fallback" yaml:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_rule#view_id TeamsRule#view_id}
	ViewId *string `field:"optional" json:"viewId" yaml:"viewId"`
}

