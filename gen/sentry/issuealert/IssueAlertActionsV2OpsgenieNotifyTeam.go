package issuealert


type IssueAlertActionsV2OpsgenieNotifyTeam struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#account IssueAlert#account}.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#priority IssueAlert#priority}.
	Priority *string `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#team IssueAlert#team}.
	Team *string `field:"required" json:"team" yaml:"team"`
}

