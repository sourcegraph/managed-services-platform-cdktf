package alert


type AlertActionFiltersActionsVsts struct {
	// The ID of the OpsGenie integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The ID of the Azure DevOps project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#project Alert#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The type of work item to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#work_item_type Alert#work_item_type}
	WorkItemType *string `field:"required" json:"workItemType" yaml:"workItemType"`
}

