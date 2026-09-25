package metricalert


type MetricAlertTriggerAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#target_type MetricAlert#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#type MetricAlert#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Slack channel ID to avoid rate-limiting, see [here](https://docs.sentry.io/product/integrations/notification-incidents/slack/#rate-limiting-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#input_channel_id MetricAlert#input_channel_id}
	InputChannelId *string `field:"optional" json:"inputChannelId" yaml:"inputChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#integration_id MetricAlert#integration_id}.
	IntegrationId *float64 `field:"optional" json:"integrationId" yaml:"integrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#priority MetricAlert#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// The ID of the Sentry App (required when type is sentry_app).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#sentry_app_id MetricAlert#sentry_app_id}
	SentryAppId *float64 `field:"optional" json:"sentryAppId" yaml:"sentryAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_alert#target_identifier MetricAlert#target_identifier}.
	TargetIdentifier *string `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

