package dialogflowcxflow


type DialogflowCxFlowAdvancedSettingsDtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequency) events.
	//
	// For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will detect the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/dialogflow_cx_flow#enabled DialogflowCxFlow#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The digit that terminates a DTMF digit sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/dialogflow_cx_flow#finish_digit DialogflowCxFlow#finish_digit}
	FinishDigit *string `field:"optional" json:"finishDigit" yaml:"finishDigit"`
	// Max length of DTMF digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.7.0/docs/resources/dialogflow_cx_flow#max_digits DialogflowCxFlow#max_digits}
	MaxDigits *float64 `field:"optional" json:"maxDigits" yaml:"maxDigits"`
}
