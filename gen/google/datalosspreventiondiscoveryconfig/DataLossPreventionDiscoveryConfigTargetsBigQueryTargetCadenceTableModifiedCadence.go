package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceTableModifiedCadence struct {
	// How frequently data profiles can be updated when tables are modified. Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#frequency DataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The type of events to consider when deciding if the table has been modified and should have the profile updated.
	//
	// Defaults to MODIFIED_TIMESTAMP Possible values: ["TABLE_MODIFIED_TIMESTAMP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#types DataLossPreventionDiscoveryConfig#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

