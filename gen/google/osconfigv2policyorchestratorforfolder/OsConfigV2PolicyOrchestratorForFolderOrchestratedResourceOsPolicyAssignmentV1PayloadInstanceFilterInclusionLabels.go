package osconfigv2policyorchestratorforfolder


type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilterInclusionLabels struct {
	// Labels are identified by key/value pairs in this map.
	//
	// A VM should contain all the key/value pairs specified in this
	// map to be selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#labels OsConfigV2PolicyOrchestratorForFolder#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

