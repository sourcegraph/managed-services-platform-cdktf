package googlepubsubsubscription


type GooglePubsubSubscriptionCloudStorageConfig struct {
	// User-provided name for the Cloud Storage bucket.
	//
	// The bucket must be created by the user. The bucket name must be without any prefix like "gs://".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#bucket GooglePubsubSubscription#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// avro_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#avro_config GooglePubsubSubscription#avro_config}
	AvroConfig *GooglePubsubSubscriptionCloudStorageConfigAvroConfig `field:"optional" json:"avroConfig" yaml:"avroConfig"`
	// User-provided format string specifying how to represent datetimes in Cloud Storage filenames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#filename_datetime_format GooglePubsubSubscription#filename_datetime_format}
	FilenameDatetimeFormat *string `field:"optional" json:"filenameDatetimeFormat" yaml:"filenameDatetimeFormat"`
	// User-provided prefix for Cloud Storage filename.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#filename_prefix GooglePubsubSubscription#filename_prefix}
	FilenamePrefix *string `field:"optional" json:"filenamePrefix" yaml:"filenamePrefix"`
	// User-provided suffix for Cloud Storage filename. Must not end in "/".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#filename_suffix GooglePubsubSubscription#filename_suffix}
	FilenameSuffix *string `field:"optional" json:"filenameSuffix" yaml:"filenameSuffix"`
	// The maximum bytes that can be written to a Cloud Storage file before a new file is created.
	//
	// Min 1 KB, max 10 GiB.
	// The maxBytes limit may be exceeded in cases where messages are larger than the limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#max_bytes GooglePubsubSubscription#max_bytes}
	MaxBytes *float64 `field:"optional" json:"maxBytes" yaml:"maxBytes"`
	// The maximum duration that can elapse before a new Cloud Storage file is created.
	//
	// Min 1 minute, max 10 minutes, default 5 minutes.
	// May not exceed the subscription's acknowledgement deadline.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#max_duration GooglePubsubSubscription#max_duration}
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
	// The maximum messages that can be written to a Cloud Storage file before a new file is created.
	//
	// Min 1000 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#max_messages GooglePubsubSubscription#max_messages}
	MaxMessages *float64 `field:"optional" json:"maxMessages" yaml:"maxMessages"`
	// The service account to use to write to Cloud Storage. If not specified, the Pub/Sub [service agent](https://cloud.google.com/iam/docs/service-agents), service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com, is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_subscription#service_account_email GooglePubsubSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

