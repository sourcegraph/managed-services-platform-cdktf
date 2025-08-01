package healthcarehl7v2store


type HealthcareHl7V2StoreNotificationConfigs struct {
	// The Cloud Pub/Sub topic that notifications of changes are published on.
	//
	// Supplied by the client.
	// PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	// It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	// was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	// project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	// Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
	//
	// If a notification cannot be published to Cloud Pub/Sub, errors will be logged to Stackdriver
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/healthcare_hl7_v2_store#pubsub_topic HealthcareHl7V2Store#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// Restricts notifications sent for messages matching a filter. If this is empty, all messages are matched. Syntax: https://cloud.google.com/appengine/docs/standard/python/search/query_strings.
	//
	// Fields/functions available for filtering are:
	//
	// * messageType, from the MSH-9.1 field. For example, NOT messageType = "ADT".
	// * send_date or sendDate, the YYYY-MM-DD date the message was sent in the dataset's timeZone, from the MSH-7 segment. For example, send_date < "2017-01-02".
	// * sendTime, the timestamp when the message was sent, using the RFC3339 time format for comparisons, from the MSH-7 segment. For example, sendTime < "2017-01-02T00:00:00-05:00".
	// * sendFacility, the care center that the message came from, from the MSH-4 segment. For example, sendFacility = "ABC".
	// * PatientId(value, type), which matches if the message lists a patient having an ID of the given value and type in the PID-2, PID-3, or PID-4 segments. For example, PatientId("123456", "MRN").
	// * labels.x, a string value of the label with key x as set using the Message.labels map. For example, labels."priority"="high". The operator :* can be used to assert the existence of a label. For example, labels."priority":*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/healthcare_hl7_v2_store#filter HealthcareHl7V2Store#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

