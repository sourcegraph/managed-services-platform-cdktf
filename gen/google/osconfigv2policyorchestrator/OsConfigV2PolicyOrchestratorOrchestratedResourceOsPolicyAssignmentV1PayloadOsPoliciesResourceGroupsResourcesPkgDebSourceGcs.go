package osconfigv2policyorchestrator


type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs struct {
	// Required. Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator#bucket OsConfigV2PolicyOrchestrator#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Required. Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator#object OsConfigV2PolicyOrchestrator#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator#generation OsConfigV2PolicyOrchestrator#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

