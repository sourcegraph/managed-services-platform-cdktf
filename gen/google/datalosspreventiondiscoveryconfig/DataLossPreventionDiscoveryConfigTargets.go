package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargets struct {
	// big_query_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_loss_prevention_discovery_config#big_query_target DataLossPreventionDiscoveryConfig#big_query_target}
	BigQueryTarget *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget `field:"optional" json:"bigQueryTarget" yaml:"bigQueryTarget"`
}

