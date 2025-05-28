package projectownership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectOwnershipConfig struct {
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
	// The auto-assignment mode.
	//
	// The options are: `none` - No auto-assignment, `all` - Assign all issues, `unhandled` - Assign unhandled issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#auto_assignment ProjectOwnership#auto_assignment}
	AutoAssignment *string `field:"required" json:"autoAssignment" yaml:"autoAssignment"`
	// Whether to automatically sync codeowners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#codeowners_auto_sync ProjectOwnership#codeowners_auto_sync}
	CodeownersAutoSync interface{} `field:"required" json:"codeownersAutoSync" yaml:"codeownersAutoSync"`
	// Whether to fall through to the default ownership rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#fallthrough ProjectOwnership#fallthrough}
	Fallthrough interface{} `field:"required" json:"fallthrough" yaml:"fallthrough"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#organization ProjectOwnership#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#project ProjectOwnership#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Raw input for ownership configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_ownership#raw ProjectOwnership#raw}
	Raw *string `field:"required" json:"raw" yaml:"raw"`
}

