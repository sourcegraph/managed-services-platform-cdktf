package metricmonitor


type MetricMonitorIssueDetection struct {
	// `static`: Threshold based monitor; `percent`: Change based monitor; `dynamic`: Dynamic monitor. Valid values are: `static`, `percent`, and `dynamic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#type MetricMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The comparison delta in seconds to use for the aggregate query. Only required for `percent` type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#comparison_delta MetricMonitor#comparison_delta}
	ComparisonDelta *float64 `field:"optional" json:"comparisonDelta" yaml:"comparisonDelta"`
}

