package notificationpolicy


type NotificationPolicyTimeRestrictionRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#end_day NotificationPolicy#end_day}.
	EndDay *string `field:"required" json:"endDay" yaml:"endDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#end_hour NotificationPolicy#end_hour}.
	EndHour *float64 `field:"required" json:"endHour" yaml:"endHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#end_min NotificationPolicy#end_min}.
	EndMin *float64 `field:"required" json:"endMin" yaml:"endMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#start_day NotificationPolicy#start_day}.
	StartDay *string `field:"required" json:"startDay" yaml:"startDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#start_hour NotificationPolicy#start_hour}.
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#start_min NotificationPolicy#start_min}.
	StartMin *float64 `field:"required" json:"startMin" yaml:"startMin"`
}

