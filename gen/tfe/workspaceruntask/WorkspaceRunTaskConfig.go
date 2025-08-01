package workspaceruntask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceRunTaskConfig struct {
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
	// The enforcement level of the task. Valid values are `advisory` and `mandatory`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/workspace_run_task#enforcement_level WorkspaceRunTask#enforcement_level}
	EnforcementLevel *string `field:"required" json:"enforcementLevel" yaml:"enforcementLevel"`
	// The id of the Run task to associate to the Workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/workspace_run_task#task_id WorkspaceRunTask#task_id}
	TaskId *string `field:"required" json:"taskId" yaml:"taskId"`
	// The id of the workspace to associate the Run task to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/workspace_run_task#workspace_id WorkspaceRunTask#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The stage to run the task in. Valid values are `pre_plan`, `post_plan`, `pre_apply` and `post_apply`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/workspace_run_task#stage WorkspaceRunTask#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The stages to run the task in. Valid values are `pre_plan`, `post_plan`, `pre_apply` and `post_apply`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/workspace_run_task#stages WorkspaceRunTask#stages}
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

