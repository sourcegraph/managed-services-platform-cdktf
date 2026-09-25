package metricmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricMonitorConfig struct {
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
	// Aggregate query to run on the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#aggregate MetricMonitor#aggregate}
	Aggregate *string `field:"required" json:"aggregate" yaml:"aggregate"`
	// Issue detection condition group configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#condition_group MetricMonitor#condition_group}
	ConditionGroup *MetricMonitorConditionGroup `field:"required" json:"conditionGroup" yaml:"conditionGroup"`
	// Dataset to run the aggregate query on.
	//
	// Valid values are: `events`, `transactions`, `discover`, `outcomes`, `outcomes_raw`, `sessions`, `metrics`, `generic_metrics`, `replays`, `profiles`, `search_issues`, `functions`, `spans`, and `events_analytics_platform`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#dataset MetricMonitor#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Event types to run the aggregate query on. Valid values are: `error`, `default`, `transaction`, `trace_item_span`, `trace_item_log`, and `trace_item_metric`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#event_types MetricMonitor#event_types}
	EventTypes *[]*string `field:"required" json:"eventTypes" yaml:"eventTypes"`
	// The issue detection type configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#issue_detection MetricMonitor#issue_detection}
	IssueDetection *MetricMonitorIssueDetection `field:"required" json:"issueDetection" yaml:"issueDetection"`
	// The name of this monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#name MetricMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization slug or internal ID to create the monitor for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#organization MetricMonitor#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project slug or internal ID to create the monitor for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#project MetricMonitor#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// A description of the monitor. Will be used in the resulting issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#description MetricMonitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the monitor is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#enabled MetricMonitor#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Environment to run the aggregate query on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#environment MetricMonitor#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Extrapolation mode to use for the aggregate query. Valid values are: `unknown`, `none`, `client_and_server_weighted`, and `server_weighted`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#extrapolation_mode MetricMonitor#extrapolation_mode}
	ExtrapolationMode *string `field:"optional" json:"extrapolationMode" yaml:"extrapolationMode"`
	// Sentry will assign new issues to this assignee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#owner MetricMonitor#owner}
	Owner *MetricMonitorOwner `field:"optional" json:"owner" yaml:"owner"`
	// An event search query to subscribe to and monitor for alerts.
	//
	// For example, to filter transactions so that only those with status code 400 are included, you could use `http.status_code:400`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#query MetricMonitor#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The type of query.
	//
	// If no value is provided, `query_type` is set to the default for the specified `dataset.` Valid values are: `error`, `performance`, and `crash_rate`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#query_type MetricMonitor#query_type}
	QueryType *string `field:"optional" json:"queryType" yaml:"queryType"`
	// The time window in seconds to use for the aggregate query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/metric_monitor#time_window_seconds MetricMonitor#time_window_seconds}
	TimeWindowSeconds *float64 `field:"optional" json:"timeWindowSeconds" yaml:"timeWindowSeconds"`
}

