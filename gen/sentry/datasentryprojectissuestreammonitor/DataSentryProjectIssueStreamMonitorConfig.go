package datasentryprojectissuestreammonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSentryProjectIssueStreamMonitorConfig struct {
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
	// The organization slug or internal ID of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/data-sources/project_issue_stream_monitor#organization DataSentryProjectIssueStreamMonitor#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project slug or internal ID of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/data-sources/project_issue_stream_monitor#project DataSentryProjectIssueStreamMonitor#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Return the first monitor found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/data-sources/project_issue_stream_monitor#first DataSentryProjectIssueStreamMonitor#first}
	First interface{} `field:"optional" json:"first" yaml:"first"`
}

