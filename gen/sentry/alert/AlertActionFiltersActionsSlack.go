package alert


type AlertActionFiltersActionsSlack struct {
	// The name of the Slack channel to send the notification to (e.g., #critical, Jane Schmidt).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#channel_name Alert#channel_name}
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The ID of the Slack integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The Slack channel ID to send the notification to.
	//
	// This is an optional field that can be used to avoid rate-limiting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#channel_id Alert#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Text to show alongside the notification.
	//
	// To @ a user, include their user id like `@<USER_ID>`. To include a clickable link, format the link and title like `<http://example.com|Click Here>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#notes Alert#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// A list of tags to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#tags Alert#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

