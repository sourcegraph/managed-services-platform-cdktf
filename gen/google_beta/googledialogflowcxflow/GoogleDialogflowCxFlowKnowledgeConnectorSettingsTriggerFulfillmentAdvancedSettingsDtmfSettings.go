package googledialogflowcxflow


type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings struct {
	// If true, incoming audio is processed for DTMF (dual tone multi frequtectency) events.
	//
	// For example, if the caller presses a button on their telephone keypad and DTMF processing is enabled, Dialogflow will de the event (e.g. a "3" was pressed) in the incoming audio and pass the event to the bot to drive business logic (e.g. when 3 is pressed, return the account balance).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_flow#enabled GoogleDialogflowCxFlow#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Endpoint timeout setting for matching dtmf input to regex.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.500s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_flow#endpointing_timeout_duration GoogleDialogflowCxFlow#endpointing_timeout_duration}
	EndpointingTimeoutDuration *string `field:"optional" json:"endpointingTimeoutDuration" yaml:"endpointingTimeoutDuration"`
	// The digit that terminates a DTMF digit sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_flow#finish_digit GoogleDialogflowCxFlow#finish_digit}
	FinishDigit *string `field:"optional" json:"finishDigit" yaml:"finishDigit"`
	// Interdigit timeout setting for matching dtmf input to regex.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.500s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_flow#interdigit_timeout_duration GoogleDialogflowCxFlow#interdigit_timeout_duration}
	InterdigitTimeoutDuration *string `field:"optional" json:"interdigitTimeoutDuration" yaml:"interdigitTimeoutDuration"`
	// Max length of DTMF digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_flow#max_digits GoogleDialogflowCxFlow#max_digits}
	MaxDigits *float64 `field:"optional" json:"maxDigits" yaml:"maxDigits"`
}

