package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#fixed GoogleOsConfigV2PolicyOrchestratorForFolder#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#percent GoogleOsConfigV2PolicyOrchestratorForFolder#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

