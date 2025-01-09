package notificationrule


type NotificationRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_rule#type NotificationRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/notification_rule#conditions NotificationRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

