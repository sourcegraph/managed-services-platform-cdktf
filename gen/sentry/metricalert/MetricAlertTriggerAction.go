package metricalert


type MetricAlertTriggerAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#target_type MetricAlert#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#type MetricAlert#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Slack channel ID to avoid rate-limiting, see [here](https://docs.sentry.io/product/integrations/notification-incidents/slack/#rate-limiting-error).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#input_channel_id MetricAlert#input_channel_id}
	InputChannelId *string `field:"optional" json:"inputChannelId" yaml:"inputChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#integration_id MetricAlert#integration_id}.
	IntegrationId *float64 `field:"optional" json:"integrationId" yaml:"integrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#target_identifier MetricAlert#target_identifier}.
	TargetIdentifier *string `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

