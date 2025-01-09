package notificationrule


type NotificationRuleSteps struct {
	// contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_rule#contact NotificationRule#contact}
	Contact interface{} `field:"required" json:"contact" yaml:"contact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_rule#enabled NotificationRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_rule#send_after NotificationRule#send_after}.
	SendAfter *float64 `field:"optional" json:"sendAfter" yaml:"sendAfter"`
}

