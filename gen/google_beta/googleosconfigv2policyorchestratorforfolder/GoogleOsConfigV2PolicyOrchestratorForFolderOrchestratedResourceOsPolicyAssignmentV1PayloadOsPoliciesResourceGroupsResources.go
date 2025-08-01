package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResources struct {
	// The id of the resource with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the OS policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#id GoogleOsConfigV2PolicyOrchestratorForFolder#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#exec GoogleOsConfigV2PolicyOrchestratorForFolder#exec}
	Exec *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec `field:"optional" json:"exec" yaml:"exec"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#file GoogleOsConfigV2PolicyOrchestratorForFolder#file}
	File *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFile `field:"optional" json:"file" yaml:"file"`
	// pkg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#pkg GoogleOsConfigV2PolicyOrchestratorForFolder#pkg}
	Pkg *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg `field:"optional" json:"pkg" yaml:"pkg"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#repository GoogleOsConfigV2PolicyOrchestratorForFolder#repository}
	Repository *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepository `field:"optional" json:"repository" yaml:"repository"`
}

