package metricmonitor


type MetricMonitorConditionGroupConditions struct {
	// When the condition is met, the result will be set to this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#condition_result MetricMonitor#condition_result}
	ConditionResult *float64 `field:"required" json:"conditionResult" yaml:"conditionResult"`
	// The type of condition.
	//
	// Valid values are: `eq`, `gte`, `gt`, `lte`, `lt`, `ne`, `anomaly_detection`, `age_comparison`, `assigned_to`, `event_attribute`, `event_created_by_detector`, `event_seen_count`, `every_event`, `existing_high_priority_issue`, `first_seen_event`, `issue_category`, `issue_occurrences`, `issue_open_duration`, `issue_priority_deescalating`, `issue_priority_equals`, `issue_priority_greater_or_equal`, `issue_resolution_change`, `issue_resolved_trigger`, `issue_type`, `latest_adopted_release`, `latest_release`, `level`, `new_high_priority_issue`, `reappeared_event`, `regression_event`, `tagged_event`, `event_frequency_count`, `event_frequency_percent`, `event_unique_user_frequency_count`, `event_unique_user_frequency_percent`, `percent_sessions_count`, `percent_sessions_percent`, and `seer_activity_trigger`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#type MetricMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The value to compare against. Only required for types other than `anomaly_detection`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#comparison MetricMonitor#comparison}
	Comparison *float64 `field:"optional" json:"comparison" yaml:"comparison"`
	// Choose your level of anomaly responsiveness.
	//
	// Higher thresholds means alerts for most anomalies. Lower thresholds means alerts only for larger ones. Only required for `anomaly_detection` type. Valid values are: `low`, `medium`, and `high`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#comparison_sensitivity MetricMonitor#comparison_sensitivity}
	ComparisonSensitivity *string `field:"optional" json:"comparisonSensitivity" yaml:"comparisonSensitivity"`
	// Decide if you want to be alerted to anomalies that are moving above, below, or in both directions in relation to your threshold.
	//
	// Only required for `anomaly_detection` type. Valid values are: `above`, `below`, and `above_and_below`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#comparison_threshold_type MetricMonitor#comparison_threshold_type}
	ComparisonThresholdType *string `field:"optional" json:"comparisonThresholdType" yaml:"comparisonThresholdType"`
}

