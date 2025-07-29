package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote struct {
	// URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#uri GoogleOsConfigV2PolicyOrchestratorForFolder#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#sha256_checksum GoogleOsConfigV2PolicyOrchestratorForFolder#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

