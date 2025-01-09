package dashboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardConfig struct {
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
	// The slug of the organization the dashboard belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/dashboard#organization Dashboard#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Dashboard title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/dashboard#widget Dashboard#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

