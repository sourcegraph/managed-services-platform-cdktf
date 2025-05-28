package issuealert


type IssueAlertActionsV2SlackNotifyService struct {
	// The name of the channel to send the notification to (e.g., #critical, Jane Schmidt).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#channel IssueAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The integration ID associated with the Slack workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#workspace IssueAlert#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Text to show alongside the notification.
	//
	// To @ a user, include their user id like `@<USER_ID>`. To include a clickable link, format the link and title like `<http://example.com|Click Here>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#notes IssueAlert#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// A string of tags to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#tags IssueAlert#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

