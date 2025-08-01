package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger struct {
	// The absolute number of time series that must fail the predicate for the condition to be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#count GoogleMonitoringAlertPolicy#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The percentage of time series that must fail the predicate for the condition to be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#percent GoogleMonitoringAlertPolicy#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

