package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargets struct {
	// big_query_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.29.0/docs/resources/google_data_loss_prevention_discovery_config#big_query_target GoogleDataLossPreventionDiscoveryConfig#big_query_target}
	BigQueryTarget *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTarget `field:"optional" json:"bigQueryTarget" yaml:"bigQueryTarget"`
}

