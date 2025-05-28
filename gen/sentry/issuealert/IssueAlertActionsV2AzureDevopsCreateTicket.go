package issuealert


type IssueAlertActionsV2AzureDevopsCreateTicket struct {
	// The integration ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#integration IssueAlert#integration}
	Integration *string `field:"required" json:"integration" yaml:"integration"`
	// The ID of the Azure DevOps project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#project IssueAlert#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The type of work item to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#work_item_type IssueAlert#work_item_type}
	WorkItemType *string `field:"required" json:"workItemType" yaml:"workItemType"`
}

