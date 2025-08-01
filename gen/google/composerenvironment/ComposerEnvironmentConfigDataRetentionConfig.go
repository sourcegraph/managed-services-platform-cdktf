package composerenvironment


type ComposerEnvironmentConfigDataRetentionConfig struct {
	// airflow_metadata_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#airflow_metadata_retention_config ComposerEnvironment#airflow_metadata_retention_config}
	AirflowMetadataRetentionConfig interface{} `field:"optional" json:"airflowMetadataRetentionConfig" yaml:"airflowMetadataRetentionConfig"`
	// task_logs_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/composer_environment#task_logs_retention_config ComposerEnvironment#task_logs_retention_config}
	TaskLogsRetentionConfig interface{} `field:"optional" json:"taskLogsRetentionConfig" yaml:"taskLogsRetentionConfig"`
}

