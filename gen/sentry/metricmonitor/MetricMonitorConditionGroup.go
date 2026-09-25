package metricmonitor


type MetricMonitorConditionGroup struct {
	// Issue detection conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#conditions MetricMonitor#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The logic to apply to the conditions.
	//
	// `any` will evaluate all conditions, and return true if any of those are met. `any-short` will stop evaluating conditions as soon as one is met. `all` will evaluate all conditions, and return true if all of those are met. `none` will return true if none of the conditions are met, will return false immediately if any are met. Valid values are: `any`, `any-short`, `all`, and `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#logic_type MetricMonitor#logic_type}
	LogicType *string `field:"optional" json:"logicType" yaml:"logicType"`
}

