package escalation


type EscalationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#condition Escalation#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#delay Escalation#delay}.
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#notify_type Escalation#notify_type}.
	NotifyType *string `field:"required" json:"notifyType" yaml:"notifyType"`
	// recipient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#recipient Escalation#recipient}
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
}

