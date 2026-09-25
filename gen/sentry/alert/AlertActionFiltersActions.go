package alert


type AlertActionFiltersActions struct {
	// Notify on Discord.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#discord Alert#discord}
	Discord *AlertActionFiltersActionsDiscord `field:"optional" json:"discord" yaml:"discord"`
	// Notify on Preferred Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#email Alert#email}
	Email *AlertActionFiltersActionsEmail `field:"optional" json:"email" yaml:"email"`
	// Create a GitHub issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#github Alert#github}
	Github *AlertActionFiltersActionsGithub `field:"optional" json:"github" yaml:"github"`
	// Create a Jira ticket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#jira Alert#jira}
	Jira *AlertActionFiltersActionsJira `field:"optional" json:"jira" yaml:"jira"`
	// Create a Jira Server ticket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#jira_server Alert#jira_server}
	JiraServer *AlertActionFiltersActionsJiraServer `field:"optional" json:"jiraServer" yaml:"jiraServer"`
	// Notify on Microsoft Teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#msteams Alert#msteams}
	Msteams *AlertActionFiltersActionsMsteams `field:"optional" json:"msteams" yaml:"msteams"`
	// Notify on OpsGenie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#opsgenie Alert#opsgenie}
	Opsgenie *AlertActionFiltersActionsOpsgenie `field:"optional" json:"opsgenie" yaml:"opsgenie"`
	// Notify on PagerDuty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#pagerduty Alert#pagerduty}
	Pagerduty *AlertActionFiltersActionsPagerduty `field:"optional" json:"pagerduty" yaml:"pagerduty"`
	// Send a notification to all legacy integrations (plugins). **Deprecated** Action type plugin is deprecated and cannot be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#plugin Alert#plugin}
	Plugin *AlertActionFiltersActionsPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	// Trigger an action in a Sentry App (e.g. Rootly).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#sentry_app Alert#sentry_app}
	SentryApp *AlertActionFiltersActionsSentryApp `field:"optional" json:"sentryApp" yaml:"sentryApp"`
	// Notify on Slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#slack Alert#slack}
	Slack *AlertActionFiltersActionsSlack `field:"optional" json:"slack" yaml:"slack"`
	// Notify on Azure DevOps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#vsts Alert#vsts}
	Vsts *AlertActionFiltersActionsVsts `field:"optional" json:"vsts" yaml:"vsts"`
	// Send a notification via a legacy integration service (e.g. an internal integration's webhook). This is the successor to the `notify_event_service` action on the legacy `sentry_issue_alert` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#webhook Alert#webhook}
	Webhook *AlertActionFiltersActionsWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

