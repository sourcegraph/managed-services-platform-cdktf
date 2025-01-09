package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigOrgConfigLocation struct {
	// The ID for the folder within an organization to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_data_loss_prevention_discovery_config#folder_id GoogleDataLossPreventionDiscoveryConfig#folder_id}
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// The ID of an organization to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_data_loss_prevention_discovery_config#organization_id GoogleDataLossPreventionDiscoveryConfig#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

