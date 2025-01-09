package workspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#name Workspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#agent_pool_id Workspace#agent_pool_id}.
	AgentPoolId *string `field:"optional" json:"agentPoolId" yaml:"agentPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#allow_destroy_plan Workspace#allow_destroy_plan}.
	AllowDestroyPlan interface{} `field:"optional" json:"allowDestroyPlan" yaml:"allowDestroyPlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#assessments_enabled Workspace#assessments_enabled}.
	AssessmentsEnabled interface{} `field:"optional" json:"assessmentsEnabled" yaml:"assessmentsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#auto_apply Workspace#auto_apply}.
	AutoApply interface{} `field:"optional" json:"autoApply" yaml:"autoApply"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#auto_apply_run_trigger Workspace#auto_apply_run_trigger}.
	AutoApplyRunTrigger interface{} `field:"optional" json:"autoApplyRunTrigger" yaml:"autoApplyRunTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#auto_destroy_activity_duration Workspace#auto_destroy_activity_duration}.
	AutoDestroyActivityDuration *string `field:"optional" json:"autoDestroyActivityDuration" yaml:"autoDestroyActivityDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#auto_destroy_at Workspace#auto_destroy_at}.
	AutoDestroyAt *string `field:"optional" json:"autoDestroyAt" yaml:"autoDestroyAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#description Workspace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#execution_mode Workspace#execution_mode}.
	ExecutionMode *string `field:"optional" json:"executionMode" yaml:"executionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#file_triggers_enabled Workspace#file_triggers_enabled}.
	FileTriggersEnabled interface{} `field:"optional" json:"fileTriggersEnabled" yaml:"fileTriggersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#force_delete Workspace#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#global_remote_state Workspace#global_remote_state}.
	GlobalRemoteState interface{} `field:"optional" json:"globalRemoteState" yaml:"globalRemoteState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#id Workspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#ignore_additional_tag_names Workspace#ignore_additional_tag_names}.
	IgnoreAdditionalTagNames interface{} `field:"optional" json:"ignoreAdditionalTagNames" yaml:"ignoreAdditionalTagNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#operations Workspace#operations}.
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#organization Workspace#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#project_id Workspace#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#queue_all_runs Workspace#queue_all_runs}.
	QueueAllRuns interface{} `field:"optional" json:"queueAllRuns" yaml:"queueAllRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#remote_state_consumer_ids Workspace#remote_state_consumer_ids}.
	RemoteStateConsumerIds *[]*string `field:"optional" json:"remoteStateConsumerIds" yaml:"remoteStateConsumerIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#source_name Workspace#source_name}.
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#source_url Workspace#source_url}.
	SourceUrl *string `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#speculative_enabled Workspace#speculative_enabled}.
	SpeculativeEnabled interface{} `field:"optional" json:"speculativeEnabled" yaml:"speculativeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#ssh_key_id Workspace#ssh_key_id}.
	SshKeyId *string `field:"optional" json:"sshKeyId" yaml:"sshKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#structured_run_output_enabled Workspace#structured_run_output_enabled}.
	StructuredRunOutputEnabled interface{} `field:"optional" json:"structuredRunOutputEnabled" yaml:"structuredRunOutputEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#tag_names Workspace#tag_names}.
	TagNames *[]*string `field:"optional" json:"tagNames" yaml:"tagNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#terraform_version Workspace#terraform_version}.
	TerraformVersion *string `field:"optional" json:"terraformVersion" yaml:"terraformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#trigger_patterns Workspace#trigger_patterns}.
	TriggerPatterns *[]*string `field:"optional" json:"triggerPatterns" yaml:"triggerPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#trigger_prefixes Workspace#trigger_prefixes}.
	TriggerPrefixes *[]*string `field:"optional" json:"triggerPrefixes" yaml:"triggerPrefixes"`
	// vcs_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#vcs_repo Workspace#vcs_repo}
	VcsRepo *WorkspaceVcsRepo `field:"optional" json:"vcsRepo" yaml:"vcsRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/workspace#working_directory Workspace#working_directory}.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

