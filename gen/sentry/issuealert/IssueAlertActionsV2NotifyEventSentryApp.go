package issuealert


type IssueAlertActionsV2NotifyEventSentryApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#sentry_app_installation_uuid IssueAlert#sentry_app_installation_uuid}.
	SentryAppInstallationUuid *string `field:"required" json:"sentryAppInstallationUuid" yaml:"sentryAppInstallationUuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#settings IssueAlert#settings}.
	Settings *map[string]*string `field:"optional" json:"settings" yaml:"settings"`
}

