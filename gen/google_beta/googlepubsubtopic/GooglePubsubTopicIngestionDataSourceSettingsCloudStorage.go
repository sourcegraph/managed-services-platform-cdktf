package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsCloudStorage struct {
	// Cloud Storage bucket. The bucket name must be without any prefix like "gs://". See the bucket naming requirements: https://cloud.google.com/storage/docs/buckets#naming.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#bucket GooglePubsubTopic#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// avro_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#avro_format GooglePubsubTopic#avro_format}
	AvroFormat *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat `field:"optional" json:"avroFormat" yaml:"avroFormat"`
	// Glob pattern used to match objects that will be ingested.
	//
	// If unset, all
	// objects will be ingested. See the supported patterns:
	// https://cloud.google.com/storage/docs/json_api/v1/objects/list#list-objects-and-prefixes-using-glob
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#match_glob GooglePubsubTopic#match_glob}
	MatchGlob *string `field:"optional" json:"matchGlob" yaml:"matchGlob"`
	// The timestamp set in RFC3339 text format.
	//
	// If set, only objects with a
	// larger or equal timestamp will be ingested. Unset by default, meaning
	// all objects will be ingested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#minimum_object_create_time GooglePubsubTopic#minimum_object_create_time}
	MinimumObjectCreateTime *string `field:"optional" json:"minimumObjectCreateTime" yaml:"minimumObjectCreateTime"`
	// pubsub_avro_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#pubsub_avro_format GooglePubsubTopic#pubsub_avro_format}
	PubsubAvroFormat *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat `field:"optional" json:"pubsubAvroFormat" yaml:"pubsubAvroFormat"`
	// text_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#text_format GooglePubsubTopic#text_format}
	TextFormat *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat `field:"optional" json:"textFormat" yaml:"textFormat"`
}

