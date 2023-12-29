package notificationpolicy


type NotificationPolicyDelayActionDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#time_amount NotificationPolicy#time_amount}.
	TimeAmount *float64 `field:"required" json:"timeAmount" yaml:"timeAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#time_unit NotificationPolicy#time_unit}.
	TimeUnit *string `field:"optional" json:"timeUnit" yaml:"timeUnit"`
}

