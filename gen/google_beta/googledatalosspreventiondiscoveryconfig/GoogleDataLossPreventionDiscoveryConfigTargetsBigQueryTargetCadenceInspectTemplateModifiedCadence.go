package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceInspectTemplateModifiedCadence struct {
	// How frequently data profiles can be updated when the template is modified.
	//
	// Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_data_loss_prevention_discovery_config#frequency GoogleDataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
}

