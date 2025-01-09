package issuealert


type IssueAlertFiltersV2IssueCategory struct {
	// Valid values are: `Error`, `Performance`, `Profile`, `Cron`, `Replay`, `Feedback`, `Uptime`, and `Metric_Alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#value IssueAlert#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

