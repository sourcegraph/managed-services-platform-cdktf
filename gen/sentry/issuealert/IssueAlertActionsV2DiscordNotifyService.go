package issuealert


type IssueAlertActionsV2DiscordNotifyService struct {
	// The ID of the channel to send the notification to.
	//
	// You must enter either a channel ID or a channel URL, not a channel name
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#channel_id IssueAlert#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The integration ID associated with the Discord server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#server IssueAlert#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// A string of tags to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#tags IssueAlert#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

