package issuealert


type IssueAlertActionsV2NotifyEventService struct {
	// The slug of the integration service. Sourced from `https://terraform-provider-sentry.sentry.io/settings/developer-settings/<service>/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#service IssueAlert#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

