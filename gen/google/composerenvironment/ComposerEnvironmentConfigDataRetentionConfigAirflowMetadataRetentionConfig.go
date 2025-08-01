package composerenvironment


type ComposerEnvironmentConfigDataRetentionConfigAirflowMetadataRetentionConfig struct {
	// How many days data should be retained for.
	//
	// This field is supported for Cloud Composer environments in composer 3 and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#retention_days ComposerEnvironment#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Whether database retention is enabled or not.
	//
	// This field is supported for Cloud Composer environments in composer 3 and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#retention_mode ComposerEnvironment#retention_mode}
	RetentionMode *string `field:"optional" json:"retentionMode" yaml:"retentionMode"`
}

