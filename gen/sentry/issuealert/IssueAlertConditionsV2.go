package issuealert


type IssueAlertConditionsV2 struct {
	// When the `comparison_type` is `count`, the number of events in an issue is more than `value` in `interval`.
	//
	// When the `comparison_type` is `percent`, the number of events in an issue is `value` % higher in `interval` compared to `comparison_interval` ago.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#event_frequency IssueAlert#event_frequency}
	EventFrequency *IssueAlertConditionsV2EventFrequency `field:"optional" json:"eventFrequency" yaml:"eventFrequency"`
	// When the `comparison_type` is `count`, the percent of sessions affected by an issue is more than `value` in `interval`.
	//
	// When the `comparison_type` is `percent`, the percent of sessions affected by an issue is `value` % higher in `interval` compared to `comparison_interval` ago.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#event_frequency_percent IssueAlert#event_frequency_percent}
	EventFrequencyPercent *IssueAlertConditionsV2EventFrequencyPercent `field:"optional" json:"eventFrequencyPercent" yaml:"eventFrequencyPercent"`
	// When the `comparison_type` is `count`, the number of users affected by an issue is more than `value` in `interval`.
	//
	// When the `comparison_type` is `percent`, the number of users affected by an issue is `value` % higher in `interval` compared to `comparison_interval` ago.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#event_unique_user_frequency IssueAlert#event_unique_user_frequency}
	EventUniqueUserFrequency *IssueAlertConditionsV2EventUniqueUserFrequency `field:"optional" json:"eventUniqueUserFrequency" yaml:"eventUniqueUserFrequency"`
	// Sentry marks an existing issue as high priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#existing_high_priority_issue IssueAlert#existing_high_priority_issue}
	ExistingHighPriorityIssue *IssueAlertConditionsV2ExistingHighPriorityIssue `field:"optional" json:"existingHighPriorityIssue" yaml:"existingHighPriorityIssue"`
	// A new issue is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#first_seen_event IssueAlert#first_seen_event}
	FirstSeenEvent *IssueAlertConditionsV2FirstSeenEvent `field:"optional" json:"firstSeenEvent" yaml:"firstSeenEvent"`
	// Sentry marks a new issue as high priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#new_high_priority_issue IssueAlert#new_high_priority_issue}
	NewHighPriorityIssue *IssueAlertConditionsV2NewHighPriorityIssue `field:"optional" json:"newHighPriorityIssue" yaml:"newHighPriorityIssue"`
	// The issue changes state from ignored to unresolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#reappeared_event IssueAlert#reappeared_event}
	ReappearedEvent *IssueAlertConditionsV2ReappearedEvent `field:"optional" json:"reappearedEvent" yaml:"reappearedEvent"`
	// The issue changes state from resolved to unresolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#regression_event IssueAlert#regression_event}
	RegressionEvent *IssueAlertConditionsV2RegressionEvent `field:"optional" json:"regressionEvent" yaml:"regressionEvent"`
}

