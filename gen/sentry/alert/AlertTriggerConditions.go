package alert


type AlertTriggerConditions struct {
	// Number of events seen by the workflow exceeds a threshold within an interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#event_frequency_count Alert#event_frequency_count}
	EventFrequencyCount *AlertTriggerConditionsEventFrequencyCount `field:"optional" json:"eventFrequencyCount" yaml:"eventFrequencyCount"`
	// A new issue is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#first_seen_event Alert#first_seen_event}
	FirstSeenEvent *AlertTriggerConditionsFirstSeenEvent `field:"optional" json:"firstSeenEvent" yaml:"firstSeenEvent"`
	// An issue is resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_resolved_trigger Alert#issue_resolved_trigger}
	IssueResolvedTrigger *AlertTriggerConditionsIssueResolvedTrigger `field:"optional" json:"issueResolvedTrigger" yaml:"issueResolvedTrigger"`
	// An issue escalates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#reappeared_event Alert#reappeared_event}
	ReappearedEvent *AlertTriggerConditionsReappearedEvent `field:"optional" json:"reappearedEvent" yaml:"reappearedEvent"`
	// A resolved issue becomes unresolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#regression_event Alert#regression_event}
	RegressionEvent *AlertTriggerConditionsRegressionEvent `field:"optional" json:"regressionEvent" yaml:"regressionEvent"`
}

