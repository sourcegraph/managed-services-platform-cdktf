package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargets struct {
	// big_query_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_discovery_config#big_query_target GoogleDataLossPreventionDiscoveryConfig#big_query_target}
	BigQueryTarget *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTarget `field:"optional" json:"bigQueryTarget" yaml:"bigQueryTarget"`
	// cloud_sql_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_discovery_config#cloud_sql_target GoogleDataLossPreventionDiscoveryConfig#cloud_sql_target}
	CloudSqlTarget *GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTarget `field:"optional" json:"cloudSqlTarget" yaml:"cloudSqlTarget"`
	// cloud_storage_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_discovery_config#cloud_storage_target GoogleDataLossPreventionDiscoveryConfig#cloud_storage_target}
	CloudStorageTarget *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTarget `field:"optional" json:"cloudStorageTarget" yaml:"cloudStorageTarget"`
	// secrets_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_discovery_config#secrets_target GoogleDataLossPreventionDiscoveryConfig#secrets_target}
	SecretsTarget *GoogleDataLossPreventionDiscoveryConfigTargetsSecretsTarget `field:"optional" json:"secretsTarget" yaml:"secretsTarget"`
}

