package issuealert


type IssueAlertFiltersV2TaggedEvent struct {
	// The tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#key IssueAlert#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The comparison operator.
	//
	// Valid values are: `CONTAINS`, `ENDS_WITH`, `EQUAL`, `GREATER_OR_EQUAL`, `GREATER`, `IS_SET`, `IS_IN`, `LESS_OR_EQUAL`, `LESS`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUAL`, `NOT_SET`, `NOT_STARTS_WITH`, `NOT_IN`, and `STARTS_WITH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#match IssueAlert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#value IssueAlert#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

