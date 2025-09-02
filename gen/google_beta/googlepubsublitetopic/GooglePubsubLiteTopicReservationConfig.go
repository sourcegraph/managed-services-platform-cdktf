package googlepubsublitetopic


type GooglePubsubLiteTopicReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_topic#throughput_reservation GooglePubsubLiteTopic#throughput_reservation}
	ThroughputReservation *string `field:"optional" json:"throughputReservation" yaml:"throughputReservation"`
}

