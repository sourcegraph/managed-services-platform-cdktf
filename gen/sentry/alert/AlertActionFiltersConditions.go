package alert


type AlertActionFiltersConditions struct {
	// Issue age.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#age_comparison Alert#age_comparison}
	AgeComparison *AlertActionFiltersConditionsAgeComparison `field:"optional" json:"ageComparison" yaml:"ageComparison"`
	// Issue assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#assigned_to Alert#assigned_to}
	AssignedTo *AlertActionFiltersConditionsAssignedTo `field:"optional" json:"assignedTo" yaml:"assignedTo"`
	// The event's `attribute` value `match` `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#event_attribute Alert#event_attribute}
	EventAttribute *AlertActionFiltersConditionsEventAttribute `field:"optional" json:"eventAttribute" yaml:"eventAttribute"`
	// Number of events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#event_frequency_count Alert#event_frequency_count}
	EventFrequencyCount *AlertActionFiltersConditionsEventFrequencyCount `field:"optional" json:"eventFrequencyCount" yaml:"eventFrequencyCount"`
	// Percent of events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#event_frequency_percent Alert#event_frequency_percent}
	EventFrequencyPercent *AlertActionFiltersConditionsEventFrequencyPercent `field:"optional" json:"eventFrequencyPercent" yaml:"eventFrequencyPercent"`
	// Number of users affected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#event_unique_user_frequency_count Alert#event_unique_user_frequency_count}
	EventUniqueUserFrequencyCount *AlertActionFiltersConditionsEventUniqueUserFrequencyCount `field:"optional" json:"eventUniqueUserFrequencyCount" yaml:"eventUniqueUserFrequencyCount"`
	// Issue category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_category Alert#issue_category}
	IssueCategory *AlertActionFiltersConditionsIssueCategory `field:"optional" json:"issueCategory" yaml:"issueCategory"`
	// Issue frequency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_occurrences Alert#issue_occurrences}
	IssueOccurrences *AlertActionFiltersConditionsIssueOccurrences `field:"optional" json:"issueOccurrences" yaml:"issueOccurrences"`
	// De-escalation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_priority_deescalating Alert#issue_priority_deescalating}
	IssuePriorityDeescalating *AlertActionFiltersConditionsIssuePriorityDeescalating `field:"optional" json:"issuePriorityDeescalating" yaml:"issuePriorityDeescalating"`
	// Issue priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_priority_greater_or_equal Alert#issue_priority_greater_or_equal}
	IssuePriorityGreaterOrEqual *AlertActionFiltersConditionsIssuePriorityGreaterOrEqual `field:"optional" json:"issuePriorityGreaterOrEqual" yaml:"issuePriorityGreaterOrEqual"`
	// Issue type is (or is not) `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_type Alert#issue_type}
	IssueType *AlertActionFiltersConditionsIssueType `field:"optional" json:"issueType" yaml:"issueType"`
	// The `release_age_type` adopted release associated with the event's issue is `age_comparison` than the latest adopted release in `environment`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#latest_adopted_release Alert#latest_adopted_release}
	LatestAdoptedRelease *AlertActionFiltersConditionsLatestAdoptedRelease `field:"optional" json:"latestAdoptedRelease" yaml:"latestAdoptedRelease"`
	// The event is from the latest release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#latest_release Alert#latest_release}
	LatestRelease *AlertActionFiltersConditionsLatestRelease `field:"optional" json:"latestRelease" yaml:"latestRelease"`
	// The event's level match `level`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#level Alert#level}
	Level *AlertActionFiltersConditionsLevel `field:"optional" json:"level" yaml:"level"`
	// Percentage of sessions affected count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#percent_sessions_count Alert#percent_sessions_count}
	PercentSessionsCount *AlertActionFiltersConditionsPercentSessionsCount `field:"optional" json:"percentSessionsCount" yaml:"percentSessionsCount"`
	// Percentage of sessions affected percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#percent_sessions_percent Alert#percent_sessions_percent}
	PercentSessionsPercent *AlertActionFiltersConditionsPercentSessionsPercent `field:"optional" json:"percentSessionsPercent" yaml:"percentSessionsPercent"`
	// The event's tags `key` match `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#tagged_event Alert#tagged_event}
	TaggedEvent *AlertActionFiltersConditionsTaggedEvent `field:"optional" json:"taggedEvent" yaml:"taggedEvent"`
}

