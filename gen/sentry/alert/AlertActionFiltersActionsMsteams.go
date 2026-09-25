package alert


type AlertActionFiltersActionsMsteams struct {
	// The name of the Microsoft Teams channel to send the notification to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#channel_name Alert#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The ID of the Microsoft Teams integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The integration ID associated with the Microsoft Teams team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#team_id Alert#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

