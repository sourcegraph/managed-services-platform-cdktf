package alert


type AlertActionFiltersActionsDiscord struct {
	// The ID of the Discord channel to send the notification to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#channel_id Alert#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The ID of the Discord integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// A list of tags to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#tags Alert#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

