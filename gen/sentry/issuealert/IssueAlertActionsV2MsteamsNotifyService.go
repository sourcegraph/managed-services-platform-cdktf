package issuealert


type IssueAlertActionsV2MsteamsNotifyService struct {
	// The name of the channel to send the notification to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#channel IssueAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The integration ID associated with the Microsoft Teams team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#team IssueAlert#team}
	Team *string `field:"required" json:"team" yaml:"team"`
}

