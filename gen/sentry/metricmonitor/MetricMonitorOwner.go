package metricmonitor


type MetricMonitorOwner struct {
	// The team internal ID to assign new issues to. Conflicts with `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#team_id MetricMonitor#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The user ID to assign new issues to. Conflicts with `team_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#user_id MetricMonitor#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

