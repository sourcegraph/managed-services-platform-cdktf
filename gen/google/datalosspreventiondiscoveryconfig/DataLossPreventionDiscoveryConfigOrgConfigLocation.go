package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigOrgConfigLocation struct {
	// The ID for the folder within an organization to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#folder_id DataLossPreventionDiscoveryConfig#folder_id}
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// The ID of an organization to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#organization_id DataLossPreventionDiscoveryConfig#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

