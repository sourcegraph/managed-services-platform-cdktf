package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#validate OsConfigV2PolicyOrchestratorForOrganization#validate}
	Validate *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#enforce OsConfigV2PolicyOrchestratorForOrganization#enforce}
	Enforce *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

