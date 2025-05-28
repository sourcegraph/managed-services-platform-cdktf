package issuealert


type IssueAlertFiltersV2AssignedTo struct {
	// Valid values are: `Unassigned`, `Team`, and `Member`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#target_type IssueAlert#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// The target's ID. Only required when `target_type` is `Team` or `Member`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#target_identifier IssueAlert#target_identifier}
	TargetIdentifier *string `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

