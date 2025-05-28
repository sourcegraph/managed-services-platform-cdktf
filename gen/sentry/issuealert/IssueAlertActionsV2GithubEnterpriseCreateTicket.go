package issuealert


type IssueAlertActionsV2GithubEnterpriseCreateTicket struct {
	// The integration ID associated with GitHub Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#integration IssueAlert#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
	// The name of the repository to create the issue in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#repo IssueAlert#repo}
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The GitHub user to assign the issue to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#assignee IssueAlert#assignee}
	Assignee *string `field:"optional" json:"assignee" yaml:"assignee"`
	// A list of labels to assign to the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#labels IssueAlert#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

