package issuealert


type IssueAlertActionsV2 struct {
	// Create an Azure DevOps work item in `integration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#azure_devops_create_ticket IssueAlert#azure_devops_create_ticket}
	AzureDevopsCreateTicket *IssueAlertActionsV2AzureDevopsCreateTicket `field:"optional" json:"azureDevopsCreateTicket" yaml:"azureDevopsCreateTicket"`
	// Send a notification to the `server` Discord server in the channel with ID or URL: `channel_id` and show tags `tags` in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#discord_notify_service IssueAlert#discord_notify_service}
	DiscordNotifyService *IssueAlertActionsV2DiscordNotifyService `field:"optional" json:"discordNotifyService" yaml:"discordNotifyService"`
	// Create a GitHub issue in `integration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#github_create_ticket IssueAlert#github_create_ticket}
	GithubCreateTicket *IssueAlertActionsV2GithubCreateTicket `field:"optional" json:"githubCreateTicket" yaml:"githubCreateTicket"`
	// Create a GitHub Enterprise issue in `integration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#github_enterprise_create_ticket IssueAlert#github_enterprise_create_ticket}
	GithubEnterpriseCreateTicket *IssueAlertActionsV2GithubEnterpriseCreateTicket `field:"optional" json:"githubEnterpriseCreateTicket" yaml:"githubEnterpriseCreateTicket"`
	// Create a Jira issue in `integration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#jira_create_ticket IssueAlert#jira_create_ticket}
	JiraCreateTicket *IssueAlertActionsV2JiraCreateTicket `field:"optional" json:"jiraCreateTicket" yaml:"jiraCreateTicket"`
	// Create a Jira Server issue in `integration`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#jira_server_create_ticket IssueAlert#jira_server_create_ticket}
	JiraServerCreateTicket *IssueAlertActionsV2JiraServerCreateTicket `field:"optional" json:"jiraServerCreateTicket" yaml:"jiraServerCreateTicket"`
	// Send a notification to the `team` Team to `channel`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#msteams_notify_service IssueAlert#msteams_notify_service}
	MsteamsNotifyService *IssueAlertActionsV2MsteamsNotifyService `field:"optional" json:"msteamsNotifyService" yaml:"msteamsNotifyService"`
	// Send a notification to `target_type` and if none can be found then send a notification to `fallthrough_type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#notify_email IssueAlert#notify_email}
	NotifyEmail *IssueAlertActionsV2NotifyEmail `field:"optional" json:"notifyEmail" yaml:"notifyEmail"`
	// Send a notification to all legacy integrations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#notify_event IssueAlert#notify_event}
	NotifyEvent *IssueAlertActionsV2NotifyEvent `field:"optional" json:"notifyEvent" yaml:"notifyEvent"`
	// Send a notification to a Sentry app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#notify_event_sentry_app IssueAlert#notify_event_sentry_app}
	NotifyEventSentryApp *IssueAlertActionsV2NotifyEventSentryApp `field:"optional" json:"notifyEventSentryApp" yaml:"notifyEventSentryApp"`
	// Send a notification via an integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#notify_event_service IssueAlert#notify_event_service}
	NotifyEventService *IssueAlertActionsV2NotifyEventService `field:"optional" json:"notifyEventService" yaml:"notifyEventService"`
	// Send a notification to Opsgenie account `account` and team `team` with `priority` priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#opsgenie_notify_team IssueAlert#opsgenie_notify_team}
	OpsgenieNotifyTeam *IssueAlertActionsV2OpsgenieNotifyTeam `field:"optional" json:"opsgenieNotifyTeam" yaml:"opsgenieNotifyTeam"`
	// Send a notification to PagerDuty account `account` and service `service` with `severity` severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#pagerduty_notify_service IssueAlert#pagerduty_notify_service}
	PagerdutyNotifyService *IssueAlertActionsV2PagerdutyNotifyService `field:"optional" json:"pagerdutyNotifyService" yaml:"pagerdutyNotifyService"`
	// Send a notification to the `workspace` Slack workspace to `channel` (optionally, an ID: `channel_id`) and show tags `tags` and notes `notes` in notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/issue_alert#slack_notify_service IssueAlert#slack_notify_service}
	SlackNotifyService *IssueAlertActionsV2SlackNotifyService `field:"optional" json:"slackNotifyService" yaml:"slackNotifyService"`
}

