package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionSqlRowCountTest struct {
	// The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value).
	//
	// The comparison is applied
	// on each time series, with the time series on
	// the left-hand side and the threshold on the
	// right-hand side. Only COMPARISON_LT and
	// COMPARISON_GT are supported currently. Possible values: ["COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#comparison GoogleMonitoringAlertPolicy#comparison}
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The value against which to compare the row count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_alert_policy#threshold GoogleMonitoringAlertPolicy#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

