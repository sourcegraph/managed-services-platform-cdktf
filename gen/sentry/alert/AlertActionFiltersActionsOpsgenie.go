package alert


type AlertActionFiltersActionsOpsgenie struct {
	// The ID of the OpsGenie integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The priority level for the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#priority Alert#priority}
	Priority *string `field:"required" json:"priority" yaml:"priority"`
	// The ID of the Opsgenie team to send the notification to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#team_id Alert#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The name of the Opsgenie team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#team_name Alert#team_name}
	TeamName *string `field:"required" json:"teamName" yaml:"teamName"`
}

