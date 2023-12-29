package notificationrule


type NotificationRuleTimeRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#type NotificationRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#restriction NotificationRule#restriction}
	Restriction *NotificationRuleTimeRestrictionRestriction `field:"optional" json:"restriction" yaml:"restriction"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#restrictions NotificationRule#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

