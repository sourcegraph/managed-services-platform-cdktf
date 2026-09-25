package uptimemonitor


type UptimeMonitorOwner struct {
	// The team internal ID to assign new issues to. Conflicts with `user_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#team_id UptimeMonitor#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The user ID to assign new issues to. Conflicts with `team_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#user_id UptimeMonitor#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

