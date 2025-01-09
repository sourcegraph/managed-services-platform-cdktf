package issuealert


type IssueAlertActionsV2PagerdutyNotifyService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#account IssueAlert#account}.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#service IssueAlert#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#severity IssueAlert#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
}

