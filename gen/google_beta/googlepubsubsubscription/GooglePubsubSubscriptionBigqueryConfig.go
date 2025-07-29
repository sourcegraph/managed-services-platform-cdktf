package googlepubsubsubscription


type GooglePubsubSubscriptionBigqueryConfig struct {
	// The name of the table to which to write data, of the form {projectId}.{datasetId}.{tableId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#table GooglePubsubSubscription#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// When true and use_topic_schema or use_table_schema is true, any fields that are a part of the topic schema or message schema that are not part of the BigQuery table schema are dropped when writing to BigQuery.
	//
	// Otherwise, the schemas must be kept in sync
	// and any messages with extra fields are not written and remain in the subscription's backlog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#drop_unknown_fields GooglePubsubSubscription#drop_unknown_fields}
	DropUnknownFields interface{} `field:"optional" json:"dropUnknownFields" yaml:"dropUnknownFields"`
	// The service account to use to write to BigQuery. If not specified, the Pub/Sub [service agent](https://cloud.google.com/iam/docs/service-agents), service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#service_account_email GooglePubsubSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// When true, use the BigQuery table's schema as the columns to write to in BigQuery.
	//
	// Messages
	// must be published in JSON format. Only one of use_topic_schema and use_table_schema can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#use_table_schema GooglePubsubSubscription#use_table_schema}
	UseTableSchema interface{} `field:"optional" json:"useTableSchema" yaml:"useTableSchema"`
	// When true, use the topic's schema as the columns to write to in BigQuery, if it exists.
	//
	// Only one of use_topic_schema and use_table_schema can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#use_topic_schema GooglePubsubSubscription#use_topic_schema}
	UseTopicSchema interface{} `field:"optional" json:"useTopicSchema" yaml:"useTopicSchema"`
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey to additional columns in the table.
	//
	// The subscription name, messageId, and publishTime fields are put in their own columns while all other message properties (other than data) are written to a JSON object in the attributes column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#write_metadata GooglePubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

