package googleeventarctrigger


type GoogleEventarcTriggerTransportPubsub struct {
	// Optional.
	//
	// The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: 'projects/{PROJECT_ID}/topics/{TOPIC_NAME}. You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished' only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_trigger#topic GoogleEventarcTrigger#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

