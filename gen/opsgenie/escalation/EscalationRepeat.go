package escalation


type EscalationRepeat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#close_alert_after_all Escalation#close_alert_after_all}.
	CloseAlertAfterAll interface{} `field:"optional" json:"closeAlertAfterAll" yaml:"closeAlertAfterAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#count Escalation#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#reset_recipient_states Escalation#reset_recipient_states}.
	ResetRecipientStates interface{} `field:"optional" json:"resetRecipientStates" yaml:"resetRecipientStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/escalation#wait_interval Escalation#wait_interval}.
	WaitInterval *float64 `field:"optional" json:"waitInterval" yaml:"waitInterval"`
}

