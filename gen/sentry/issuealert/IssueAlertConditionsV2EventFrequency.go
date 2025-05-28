package issuealert


type IssueAlertConditionsV2EventFrequency struct {
	// Valid values are: `count`, and `percent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#comparison_type IssueAlert#comparison_type}
	ComparisonType *string `field:"required" json:"comparisonType" yaml:"comparisonType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#value IssueAlert#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// `m` for minutes, `h` for hours, `d` for days, and `w` for weeks.
	//
	// Valid values are: `5m`, `15m`, `1h`, `1d`, `1w`, and `30d`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#comparison_interval IssueAlert#comparison_interval}
	ComparisonInterval *string `field:"optional" json:"comparisonInterval" yaml:"comparisonInterval"`
	// `m` for minutes, `h` for hours, `d` for days, and `w` for weeks.
	//
	// Valid values are: `1m`, `5m`, `15m`, `1h`, `1d`, `1w`, and `30d`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#interval IssueAlert#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
}

