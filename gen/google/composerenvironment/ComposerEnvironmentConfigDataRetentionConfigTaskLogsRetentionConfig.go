package composerenvironment


type ComposerEnvironmentConfigDataRetentionConfigTaskLogsRetentionConfig struct {
	// Whether logs in cloud logging only is enabled or not.
	//
	// This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4 and newer but not in composer-3*
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#storage_mode ComposerEnvironment#storage_mode}
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
}

