package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionEventTrigger struct {
	// event_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#event_filters GoogleCloudfunctions2Function#event_filters}
	EventFilters interface{} `field:"optional" json:"eventFilters" yaml:"eventFilters"`
	// Required. The type of event to observe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#event_type GoogleCloudfunctions2Function#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The name of a Pub/Sub topic in the same project that will be used as the transport topic for the event delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#pubsub_topic GoogleCloudfunctions2Function#pubsub_topic}
	PubsubTopic *string `field:"optional" json:"pubsubTopic" yaml:"pubsubTopic"`
	// Describes the retry policy in case of function's execution failure.
	//
	// Retried execution is charged as any other execution. Possible values: ["RETRY_POLICY_UNSPECIFIED", "RETRY_POLICY_DO_NOT_RETRY", "RETRY_POLICY_RETRY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#retry_policy GoogleCloudfunctions2Function#retry_policy}
	RetryPolicy *string `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Optional.
	//
	// The email of the trigger's service account. The service account
	// must have permission to invoke Cloud Run services. If empty, defaults to the
	// Compute Engine default service account: {project_number}-compute@developer.gserviceaccount.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#service_account_email GoogleCloudfunctions2Function#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The region that the trigger will be in.
	//
	// The trigger will only receive
	// events originating in this region. It can be the same
	// region as the function, a different region or multi-region, or the global
	// region. If not provided, defaults to the same region as the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#trigger_region GoogleCloudfunctions2Function#trigger_region}
	TriggerRegion *string `field:"optional" json:"triggerRegion" yaml:"triggerRegion"`
}

