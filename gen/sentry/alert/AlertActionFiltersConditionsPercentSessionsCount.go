package alert


type AlertActionFiltersConditionsPercentSessionsCount struct {
	// The time period in which to evaluate the value.
	//
	// e.g. Percentage of sessions affected by an issue is more than `value` in `interval`. Valid values are: `1m`, `5m`, `15m`, `1h`, `1d`, `1w`, and `30d`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#interval Alert#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// A positive integer representing the number of events in an issue that must come in before the alert will fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

