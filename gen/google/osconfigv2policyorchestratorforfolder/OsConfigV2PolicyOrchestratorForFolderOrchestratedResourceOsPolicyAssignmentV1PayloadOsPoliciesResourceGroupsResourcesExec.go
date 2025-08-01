package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#validate OsConfigV2PolicyOrchestratorForFolder#validate}
	Validate *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#enforce OsConfigV2PolicyOrchestratorForFolder#enforce}
	Enforce *OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

