package issuealert


type IssueAlertFiltersV2IssueCategory struct {
	// Valid values are: `Error`, `Performance`, `Profile`, `Cron`, `Replay`, `Feedback`, `Uptime`, `Metric_Alert`, `Test_Notification`, `Outage`, `Performance_Regression`, `User_Experience`, `Responsiveness`, and `Performance_Best_Practice`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#value IssueAlert#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

