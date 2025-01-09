package datasentryallkeys

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSentryAllKeysConfig struct {
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
	// The organization the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/data-sources/all_keys#organization DataSentryAllKeys#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/data-sources/all_keys#project DataSentryAllKeys#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Filter client keys by `active` or `inactive`. Defaults to returning all keys if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/data-sources/all_keys#filter_status DataSentryAllKeys#filter_status}
	FilterStatus *string `field:"optional" json:"filterStatus" yaml:"filterStatus"`
}

