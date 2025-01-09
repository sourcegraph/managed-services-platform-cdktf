package notificationpolicy


type NotificationPolicyTimeRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#type NotificationPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#restriction NotificationPolicy#restriction}
	Restriction *NotificationPolicyTimeRestrictionRestriction `field:"optional" json:"restriction" yaml:"restriction"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_policy#restrictions NotificationPolicy#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

