package issuealert


type IssueAlertFiltersV2LatestAdoptedRelease struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#environment IssueAlert#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Valid values are: `older`, and `newer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#older_or_newer IssueAlert#older_or_newer}
	OlderOrNewer *string `field:"required" json:"olderOrNewer" yaml:"olderOrNewer"`
	// Valid values are: `oldest`, and `newest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#oldest_or_newest IssueAlert#oldest_or_newest}
	OldestOrNewest *string `field:"required" json:"oldestOrNewest" yaml:"oldestOrNewest"`
}

