package notificationrule


type NotificationRuleStepsContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#method NotificationRule#method}.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#to NotificationRule#to}.
	To *string `field:"required" json:"to" yaml:"to"`
}

