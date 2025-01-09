package issuealert


type IssueAlertFiltersV2 struct {
	// The issue is older or newer than `value` `time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#age_comparison IssueAlert#age_comparison}
	AgeComparison *IssueAlertFiltersV2AgeComparison `field:"optional" json:"ageComparison" yaml:"ageComparison"`
	// The issue is assigned to no one, team, or member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#assigned_to IssueAlert#assigned_to}
	AssignedTo *IssueAlertFiltersV2AssignedTo `field:"optional" json:"assignedTo" yaml:"assignedTo"`
	// The event's `attribute` value `match` `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#event_attribute IssueAlert#event_attribute}
	EventAttribute *IssueAlertFiltersV2EventAttribute `field:"optional" json:"eventAttribute" yaml:"eventAttribute"`
	// The issue's category is equal to `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#issue_category IssueAlert#issue_category}
	IssueCategory *IssueAlertFiltersV2IssueCategory `field:"optional" json:"issueCategory" yaml:"issueCategory"`
	// The issue has happened at least `value` times (Note: this is approximate).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#issue_occurrences IssueAlert#issue_occurrences}
	IssueOccurrences *IssueAlertFiltersV2IssueOccurrences `field:"optional" json:"issueOccurrences" yaml:"issueOccurrences"`
	// The {oldest_or_newest} adopted release associated with the event's issue is {older_or_newer} than the latest adopted release in {environment}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#latest_adopted_release IssueAlert#latest_adopted_release}
	LatestAdoptedRelease *IssueAlertFiltersV2LatestAdoptedRelease `field:"optional" json:"latestAdoptedRelease" yaml:"latestAdoptedRelease"`
	// The event is from the latest release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#latest_release IssueAlert#latest_release}
	LatestRelease *IssueAlertFiltersV2LatestRelease `field:"optional" json:"latestRelease" yaml:"latestRelease"`
	// The event's level is `match` `level`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#level IssueAlert#level}
	Level *IssueAlertFiltersV2Level `field:"optional" json:"level" yaml:"level"`
	// The event's tags match `key` `match` `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#tagged_event IssueAlert#tagged_event}
	TaggedEvent *IssueAlertFiltersV2TaggedEvent `field:"optional" json:"taggedEvent" yaml:"taggedEvent"`
}

