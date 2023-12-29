package notificationrule


type NotificationRuleSchedules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#name NotificationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#type NotificationRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

