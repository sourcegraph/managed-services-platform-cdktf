package workspacesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/workspace_settings#workspace_id WorkspaceSettings#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/workspace_settings#agent_pool_id WorkspaceSettings#agent_pool_id}.
	AgentPoolId *string `field:"optional" json:"agentPoolId" yaml:"agentPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/workspace_settings#execution_mode WorkspaceSettings#execution_mode}.
	ExecutionMode *string `field:"optional" json:"executionMode" yaml:"executionMode"`
}

