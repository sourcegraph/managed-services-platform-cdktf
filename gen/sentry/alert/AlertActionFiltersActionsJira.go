package alert


type AlertActionFiltersActionsJira struct {
	// The ID of the Jira integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The ID of the type of issue that the ticket should be created as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#issue_type Alert#issue_type}
	IssueType *string `field:"required" json:"issueType" yaml:"issueType"`
	// The ID of the Jira project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#project Alert#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Additional Jira fields to set on the created issue, keyed by Jira field ID (e.g. `customfield_10101`). Use this for custom fields and any built-in field not exposed above. Sentry's API converts camelCase keys to snake_case on write, which corrupts them, so camelCase field IDs must be written all-lowercase: use `fixversions`, not `fixVersions`. Jira matches field IDs case-insensitively, so the lowercase spelling resolves to the same field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#additional_fields Alert#additional_fields}
	AdditionalFields *map[string]*string `field:"optional" json:"additionalFields" yaml:"additionalFields"`
	// The IDs of the Jira components to assign to the issue, used for triage routing.
	//
	// These are component IDs, not names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#components Alert#components}
	Components *[]*string `field:"optional" json:"components" yaml:"components"`
	// A comma-separated list of labels to add to the issue (e.g. `oncall,triage`). Note: unlike the `github` action's `labels`, Jira expects a single comma-separated string rather than a list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#labels Alert#labels}
	Labels *string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the priority to set on the issue. This is a priority ID, not a name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#priority Alert#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// The Jira account ID of the user to set as the reporter of the issue.
	//
	// Useful for attributing automated tickets to a service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#reporter Alert#reporter}
	Reporter *string `field:"optional" json:"reporter" yaml:"reporter"`
}

