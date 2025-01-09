package dataopsgenieescalation


type DataOpsgenieEscalationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/data-sources/escalation#condition DataOpsgenieEscalation#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/data-sources/escalation#delay DataOpsgenieEscalation#delay}.
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/data-sources/escalation#notify_type DataOpsgenieEscalation#notify_type}.
	NotifyType *string `field:"required" json:"notifyType" yaml:"notifyType"`
	// recipient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/data-sources/escalation#recipient DataOpsgenieEscalation#recipient}
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
}

