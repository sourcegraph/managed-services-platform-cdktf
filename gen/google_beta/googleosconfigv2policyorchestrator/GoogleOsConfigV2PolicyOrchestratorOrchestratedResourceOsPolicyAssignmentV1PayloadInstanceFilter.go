package googleosconfigv2policyorchestrator

type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator#all GoogleOsConfigV2PolicyOrchestrator#all}
	All any `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator#exclusion_labels GoogleOsConfigV2PolicyOrchestrator#exclusion_labels}
	ExclusionLabels any `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator#inclusion_labels GoogleOsConfigV2PolicyOrchestrator#inclusion_labels}
	InclusionLabels any `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator#inventories GoogleOsConfigV2PolicyOrchestrator#inventories}
	Inventories any `field:"optional" json:"inventories" yaml:"inventories"`
}
