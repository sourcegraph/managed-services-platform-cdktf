package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetGenerationCadenceInspectTemplateModifiedCadence struct {
	// How frequently data profiles can be updated when the template is modified.
	//
	// Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#frequency DataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
}

