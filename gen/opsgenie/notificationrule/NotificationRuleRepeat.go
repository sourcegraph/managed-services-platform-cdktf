package notificationrule


type NotificationRuleRepeat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#loop_after NotificationRule#loop_after}.
	LoopAfter *float64 `field:"required" json:"loopAfter" yaml:"loopAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_rule#enabled NotificationRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

