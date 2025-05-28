package issuealert


type IssueAlertFiltersV2AgeComparison struct {
	// Valid values are: `older`, and `newer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#comparison_type IssueAlert#comparison_type}
	ComparisonType *string `field:"required" json:"comparisonType" yaml:"comparisonType"`
	// Valid values are: `minute`, `hour`, `day`, and `week`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#time IssueAlert#time}
	Time *string `field:"required" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#value IssueAlert#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

