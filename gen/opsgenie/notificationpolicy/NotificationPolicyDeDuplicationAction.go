package notificationpolicy


type NotificationPolicyDeDuplicationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#count NotificationPolicy#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#de_duplication_action_type NotificationPolicy#de_duplication_action_type}.
	DeDuplicationActionType *string `field:"required" json:"deDuplicationActionType" yaml:"deDuplicationActionType"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#duration NotificationPolicy#duration}
	Duration interface{} `field:"optional" json:"duration" yaml:"duration"`
}

