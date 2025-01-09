package notificationpolicy


type NotificationPolicyAutoRestartAction struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#duration NotificationPolicy#duration}
	Duration interface{} `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#max_repeat_count NotificationPolicy#max_repeat_count}.
	MaxRepeatCount *float64 `field:"required" json:"maxRepeatCount" yaml:"maxRepeatCount"`
}

