package dataopsgenieescalation


type DataOpsgenieEscalationRepeat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/escalation#close_alert_after_all DataOpsgenieEscalation#close_alert_after_all}.
	CloseAlertAfterAll interface{} `field:"optional" json:"closeAlertAfterAll" yaml:"closeAlertAfterAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/escalation#count DataOpsgenieEscalation#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/escalation#reset_recipient_states DataOpsgenieEscalation#reset_recipient_states}.
	ResetRecipientStates interface{} `field:"optional" json:"resetRecipientStates" yaml:"resetRecipientStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/escalation#wait_interval DataOpsgenieEscalation#wait_interval}.
	WaitInterval *float64 `field:"optional" json:"waitInterval" yaml:"waitInterval"`
}

