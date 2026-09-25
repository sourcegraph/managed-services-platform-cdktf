package cronmonitor


type CronMonitorSchedule struct {
	// Use the crontab syntax (e.g. `0 0 * * *`). Conflicts with `interval_value` and `interval_unit`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/cron_monitor#crontab CronMonitor#crontab}
	Crontab *string `field:"optional" json:"crontab" yaml:"crontab"`
	// Interval unit.
	//
	// Conflicts with `crontab`. Must be provided with `interval_value`. Valid values are: `year`, `month`, `week`, `day`, `hour`, and `minute`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/cron_monitor#interval_unit CronMonitor#interval_unit}
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Interval value. Conflicts with `crontab`. Must be provided with `interval_unit`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/cron_monitor#interval_value CronMonitor#interval_value}
	IntervalValue *float64 `field:"optional" json:"intervalValue" yaml:"intervalValue"`
}

