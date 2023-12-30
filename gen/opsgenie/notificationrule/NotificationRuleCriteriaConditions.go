package notificationrule


type NotificationRuleCriteriaConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#field NotificationRule#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#operation NotificationRule#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#expected_value NotificationRule#expected_value}.
	ExpectedValue *string `field:"optional" json:"expectedValue" yaml:"expectedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#key NotificationRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#not NotificationRule#not}.
	Not interface{} `field:"optional" json:"not" yaml:"not"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#order NotificationRule#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

