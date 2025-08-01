package googlecloudschedulerjob


type GoogleCloudSchedulerJobRetryConfig struct {
	// The maximum amount of time to wait before retrying a job after it fails.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#max_backoff_duration GoogleCloudSchedulerJob#max_backoff_duration}
	MaxBackoffDuration *string `field:"optional" json:"maxBackoffDuration" yaml:"maxBackoffDuration"`
	// The time between retries will double maxDoublings times.
	//
	// A job's retry interval starts at minBackoffDuration,
	// then doubles maxDoublings times, then increases linearly,
	// and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#max_doublings GoogleCloudSchedulerJob#max_doublings}
	MaxDoublings *float64 `field:"optional" json:"maxDoublings" yaml:"maxDoublings"`
	// The time limit for retrying a failed job, measured from time when an execution was first attempted.
	//
	// If specified with retryCount, the job will be retried until both limits are reached.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#max_retry_duration GoogleCloudSchedulerJob#max_retry_duration}
	MaxRetryDuration *string `field:"optional" json:"maxRetryDuration" yaml:"maxRetryDuration"`
	// The minimum amount of time to wait before retrying a job after it fails.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#min_backoff_duration GoogleCloudSchedulerJob#min_backoff_duration}
	MinBackoffDuration *string `field:"optional" json:"minBackoffDuration" yaml:"minBackoffDuration"`
	// The number of attempts that the system will make to run a job using the exponential backoff procedure described by maxDoublings.
	//
	// Values greater than 5 and negative values are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job#retry_count GoogleCloudSchedulerJob#retry_count}
	RetryCount *float64 `field:"optional" json:"retryCount" yaml:"retryCount"`
}

