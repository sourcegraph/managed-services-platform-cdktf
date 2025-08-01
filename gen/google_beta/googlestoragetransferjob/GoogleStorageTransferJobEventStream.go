package googlestoragetransferjob


type GoogleStorageTransferJobEventStream struct {
	// Specifies a unique name of the resource such as AWS SQS ARN in the form 'arn:aws:sqs:region:account_id:queue_name', or Pub/Sub subscription resource name in the form 'projects/{project}/subscriptions/{sub}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_transfer_job#name GoogleStorageTransferJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the data and time at which Storage Transfer Service stops listening for events from this stream.
	//
	// After this time, any transfers in progress will complete, but no new transfers are initiated
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_transfer_job#event_stream_expiration_time GoogleStorageTransferJob#event_stream_expiration_time}
	EventStreamExpirationTime *string `field:"optional" json:"eventStreamExpirationTime" yaml:"eventStreamExpirationTime"`
	// Specifies the date and time that Storage Transfer Service starts listening for events from this stream.
	//
	// If no start time is specified or start time is in the past, Storage Transfer Service starts listening immediately
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_transfer_job#event_stream_start_time GoogleStorageTransferJob#event_stream_start_time}
	EventStreamStartTime *string `field:"optional" json:"eventStreamStartTime" yaml:"eventStreamStartTime"`
}

