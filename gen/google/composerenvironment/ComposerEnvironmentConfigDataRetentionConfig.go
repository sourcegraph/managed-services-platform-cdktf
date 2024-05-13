package composerenvironment


type ComposerEnvironmentConfigDataRetentionConfig struct {
	// task_logs_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/composer_environment#task_logs_retention_config ComposerEnvironment#task_logs_retention_config}
	TaskLogsRetentionConfig interface{} `field:"required" json:"taskLogsRetentionConfig" yaml:"taskLogsRetentionConfig"`
}

