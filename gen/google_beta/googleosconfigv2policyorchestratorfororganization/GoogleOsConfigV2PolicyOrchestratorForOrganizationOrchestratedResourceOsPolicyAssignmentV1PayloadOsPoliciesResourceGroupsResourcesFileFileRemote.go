package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFileFileRemote struct {
	// Required. URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#uri GoogleOsConfigV2PolicyOrchestratorForOrganization#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#sha256_checksum GoogleOsConfigV2PolicyOrchestratorForOrganization#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

