package metricalert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricAlertConfig struct {
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
	// The aggregation criteria to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#aggregate MetricAlert#aggregate}
	Aggregate *string `field:"required" json:"aggregate" yaml:"aggregate"`
	// The metric alert name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#name MetricAlert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The slug of the organization the metric alert belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#organization MetricAlert#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slug of the project to create the metric alert for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#project MetricAlert#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The query filter to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#query MetricAlert#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The type of threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#threshold_type MetricAlert#threshold_type}
	ThresholdType *float64 `field:"required" json:"thresholdType" yaml:"thresholdType"`
	// The period to evaluate the Alert rule in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#time_window MetricAlert#time_window}
	TimeWindow *float64 `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#trigger MetricAlert#trigger}
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// The Sentry Alert category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#dataset MetricAlert#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// Perform Alert rule in a specific environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#environment MetricAlert#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The events type of dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#event_types MetricAlert#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#id MetricAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the owner id of this Alert rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#owner MetricAlert#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The value at which the Alert rule resolves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/metric_alert#resolve_threshold MetricAlert#resolve_threshold}
	ResolveThreshold *float64 `field:"optional" json:"resolveThreshold" yaml:"resolveThreshold"`
}

