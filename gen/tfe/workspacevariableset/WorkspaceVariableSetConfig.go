package workspacevariableset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceVariableSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace_variable_set#variable_set_id WorkspaceVariableSet#variable_set_id}.
	VariableSetId *string `field:"required" json:"variableSetId" yaml:"variableSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace_variable_set#workspace_id WorkspaceVariableSet#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace_variable_set#id WorkspaceVariableSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

