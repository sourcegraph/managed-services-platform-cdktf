package team

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamConfig struct {
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
	// The name of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team#name Team#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The slug of the organization the team should be created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team#organization Team#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The optional slug for this team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team#slug Team#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

