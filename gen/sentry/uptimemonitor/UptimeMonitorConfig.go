package uptimemonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UptimeMonitorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the environment to create uptime issues in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#environment UptimeMonitor#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The amount of time between each uptime check request. Valid values are: `60`, `300`, `600`, `1200`, `1800`, and `3600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#interval_seconds UptimeMonitor#interval_seconds}
	IntervalSeconds *float64 `field:"required" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The HTTP method to use for the request. Valid values are: `GET`, `POST`, `HEAD`, `PUT`, `DELETE`, `PATCH`, and `OPTIONS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#method UptimeMonitor#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// The name of this monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#name UptimeMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization slug or internal ID to create the monitor for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#organization UptimeMonitor#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project slug or internal ID to create the monitor for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#project UptimeMonitor#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The request timeout in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#timeout_ms UptimeMonitor#timeout_ms}
	TimeoutMs *float64 `field:"required" json:"timeoutMs" yaml:"timeoutMs"`
	// The URL to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#url UptimeMonitor#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Define conditions that must be met for the check to be considered successful.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#assertion_json UptimeMonitor#assertion_json}
	AssertionJson *string `field:"optional" json:"assertionJson" yaml:"assertionJson"`
	// The request body to send. Only applicable for methods that support a body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#body UptimeMonitor#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// A description of the monitor. Will be used in the resulting issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#description UptimeMonitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Number of consecutive failed checks required to mark monitor as down. Defaults to `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#downtime_threshold UptimeMonitor#downtime_threshold}
	DowntimeThreshold *float64 `field:"optional" json:"downtimeThreshold" yaml:"downtimeThreshold"`
	// Whether the monitor is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#enabled UptimeMonitor#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The headers to send with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#headers UptimeMonitor#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Sentry will assign new issues to this assignee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#owner UptimeMonitor#owner}
	Owner *UptimeMonitorOwner `field:"optional" json:"owner" yaml:"owner"`
	// Number of consecutive successful checks required to mark monitor as recovered. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/uptime_monitor#recovery_threshold UptimeMonitor#recovery_threshold}
	RecoveryThreshold *float64 `field:"optional" json:"recoveryThreshold" yaml:"recoveryThreshold"`
}

