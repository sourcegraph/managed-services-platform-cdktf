package issuealert


type IssueAlertActionsV2NotifyEmail struct {
	// Valid values are: `IssueOwners`, `Team`, and `Member`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#target_type IssueAlert#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// Who the notification should be sent to if there are no suggested assignees.
	//
	// Valid values are: `AllMembers`, `ActiveMembers`, and `NoOne`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#fallthrough_type IssueAlert#fallthrough_type}
	FallthroughType *string `field:"optional" json:"fallthroughType" yaml:"fallthroughType"`
	// The ID of the Member or Team the notification should be sent to.
	//
	// Only required when `target_type` is `Team` or `Member`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#target_identifier IssueAlert#target_identifier}
	TargetIdentifier *string `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

