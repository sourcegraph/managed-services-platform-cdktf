package googleosconfigv2policyorchestratorforfolder

type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#resources GoogleOsConfigV2PolicyOrchestratorForFolder#resources}
	Resources any `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#inventory_filters GoogleOsConfigV2PolicyOrchestratorForFolder#inventory_filters}
	InventoryFilters any `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}
