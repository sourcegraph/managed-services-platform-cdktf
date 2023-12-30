package notificationpolicy


type NotificationPolicyDelayAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#delay_option NotificationPolicy#delay_option}.
	DelayOption *string `field:"required" json:"delayOption" yaml:"delayOption"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#duration NotificationPolicy#duration}
	Duration interface{} `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#until_hour NotificationPolicy#until_hour}.
	UntilHour *float64 `field:"optional" json:"untilHour" yaml:"untilHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#until_minute NotificationPolicy#until_minute}.
	UntilMinute *float64 `field:"optional" json:"untilMinute" yaml:"untilMinute"`
}

