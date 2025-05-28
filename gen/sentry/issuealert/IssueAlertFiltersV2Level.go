package issuealert


type IssueAlertFiltersV2Level struct {
	// Valid values are: `sample`, `debug`, `info`, `warning`, `error`, and `fatal`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#level IssueAlert#level}
	Level *string `field:"required" json:"level" yaml:"level"`
	// The comparison operator. Valid values are: `EQUAL`, `GREATER_OR_EQUAL`, and `LESS_OR_EQUAL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#match IssueAlert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
}

