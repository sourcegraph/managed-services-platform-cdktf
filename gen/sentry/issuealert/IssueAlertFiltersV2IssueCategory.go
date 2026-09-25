package issuealert


type IssueAlertFiltersV2IssueCategory struct {
	// Valid values are: `Error`, `Performance`, `Profile`, `Cron`, `Replay`, `Feedback`, `Uptime`, `Metric_Alert`, `Test_Notification`, `Outage`, `Metric`, `Db_Query`, `Http_Client`, `Frontend`, `Mobile`, `Ai_Detected`, `Preprod`, and `Configuration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/issue_alert#value IssueAlert#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

