package metricalert


type MetricAlertTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/metric_alert#alert_threshold MetricAlert#alert_threshold}.
	AlertThreshold *float64 `field:"required" json:"alertThreshold" yaml:"alertThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/metric_alert#label MetricAlert#label}.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/metric_alert#threshold_type MetricAlert#threshold_type}.
	ThresholdType *float64 `field:"required" json:"thresholdType" yaml:"thresholdType"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/metric_alert#action MetricAlert#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.3/docs/resources/metric_alert#resolve_threshold MetricAlert#resolve_threshold}.
	ResolveThreshold *float64 `field:"optional" json:"resolveThreshold" yaml:"resolveThreshold"`
}

