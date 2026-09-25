package alert


type AlertTriggerConditionsEventFrequencyCount struct {
	// The time period in which to evaluate the event count.
	//
	// Valid values are: `1m`, `5m`, `15m`, `1h`, `1d`, `1w`, and `30d`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#interval Alert#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The number of events that must be exceeded before the alert will fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

