package stack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackConfig struct {
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
	// Name of the Stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#name Stack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the project that the Stack belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#project_id Stack#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Description of the Stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#description Stack#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// vcs_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#vcs_repo Stack#vcs_repo}
	VcsRepo *StackVcsRepo `field:"optional" json:"vcsRepo" yaml:"vcsRepo"`
}

