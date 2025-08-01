package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionPrometheusQueryLanguage struct {
	// The PromQL expression to evaluate.
	//
	// Every evaluation cycle this
	// expression is evaluated at the current time, and all resultant time
	// series become pending/firing alerts. This field must not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#query GoogleMonitoringAlertPolicy#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The alerting rule name of this alert in the corresponding Prometheus configuration file.
	//
	// Some external tools may require this field to be populated correctly
	// in order to refer to the original Prometheus configuration file.
	// The rule group name and the alert name are necessary to update the
	// relevant AlertPolicies in case the definition of the rule group changes
	// in the future.
	//
	// This field is optional. If this field is not empty, then it must be a
	// valid Prometheus label name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#alert_rule GoogleMonitoringAlertPolicy#alert_rule}
	AlertRule *string `field:"optional" json:"alertRule" yaml:"alertRule"`
	// Whether to disable metric existence validation for this condition.
	//
	// This allows alerting policies to be defined on metrics that do not yet
	// exist, improving advanced customer workflows such as configuring
	// alerting policies using Terraform.
	//
	// Users with the 'monitoring.alertPolicyViewer' role are able to see the
	// name of the non-existent metric in the alerting policy condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#disable_metric_validation GoogleMonitoringAlertPolicy#disable_metric_validation}
	DisableMetricValidation interface{} `field:"optional" json:"disableMetricValidation" yaml:"disableMetricValidation"`
	// Alerts are considered firing once their PromQL expression evaluated to be "true" for this long.
	//
	// Alerts whose PromQL expression was not
	// evaluated to be "true" for long enough are considered pending. The
	// default value is zero. Must be zero or positive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#duration GoogleMonitoringAlertPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// How often this rule should be evaluated.
	//
	// Must be a positive multiple
	// of 30 seconds or missing. The default value is 30 seconds. If this
	// PrometheusQueryLanguageCondition was generated from a Prometheus
	// alerting rule, then this value should be taken from the enclosing
	// rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#evaluation_interval GoogleMonitoringAlertPolicy#evaluation_interval}
	EvaluationInterval *string `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
	// Labels to add to or overwrite in the PromQL query result. Label names must be valid.
	//
	// Label values can be templatized by using variables. The only available
	// variable names are the names of the labels in the PromQL result,
	// although label names beginning with \_\_ (two "\_") are reserved for
	// internal use. "labels" may be empty. This field is intended to be used
	// for organizing and identifying the AlertPolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#labels GoogleMonitoringAlertPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The rule group name of this alert in the corresponding Prometheus configuration file.
	//
	// Some external tools may require this field to be populated correctly
	// in order to refer to the original Prometheus configuration file.
	// The rule group name and the alert name are necessary to update the
	// relevant AlertPolicies in case the definition of the rule group changes
	// in the future. This field is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#rule_group GoogleMonitoringAlertPolicy#rule_group}
	RuleGroup *string `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
}

