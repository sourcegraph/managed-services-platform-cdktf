package issuealert


type IssueAlertActionsV2JiraCreateTicket struct {
	// The integration ID associated with Jira.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#integration IssueAlert#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
	// The ID of the type of issue that the ticket should be created as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#issue_type IssueAlert#issue_type}
	IssueType *string `field:"required" json:"issueType" yaml:"issueType"`
	// The ID of the Jira project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#project IssueAlert#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

