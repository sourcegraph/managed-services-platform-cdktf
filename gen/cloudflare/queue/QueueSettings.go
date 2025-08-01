package queue


type QueueSettings struct {
	// Number of seconds to delay delivery of all messages to consumers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/queue#delivery_delay Queue#delivery_delay}
	DeliveryDelay *float64 `field:"optional" json:"deliveryDelay" yaml:"deliveryDelay"`
	// Indicates if message delivery to consumers is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/queue#delivery_paused Queue#delivery_paused}
	DeliveryPaused interface{} `field:"optional" json:"deliveryPaused" yaml:"deliveryPaused"`
	// Number of seconds after which an unconsumed message will be delayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/queue#message_retention_period Queue#message_retention_period}
	MessageRetentionPeriod *float64 `field:"optional" json:"messageRetentionPeriod" yaml:"messageRetentionPeriod"`
}

