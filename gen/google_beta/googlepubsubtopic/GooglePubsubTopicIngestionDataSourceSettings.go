package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettings struct {
	// aws_kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_pubsub_topic#aws_kinesis GooglePubsubTopic#aws_kinesis}
	AwsKinesis *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis `field:"optional" json:"awsKinesis" yaml:"awsKinesis"`
	// cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_pubsub_topic#cloud_storage GooglePubsubTopic#cloud_storage}
	CloudStorage *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage `field:"optional" json:"cloudStorage" yaml:"cloudStorage"`
	// platform_logs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_pubsub_topic#platform_logs_settings GooglePubsubTopic#platform_logs_settings}
	PlatformLogsSettings *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings `field:"optional" json:"platformLogsSettings" yaml:"platformLogsSettings"`
}

