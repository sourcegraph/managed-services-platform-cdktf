package googledialogflowcxpage


type GoogleDialogflowCxPageEntryFulfillmentMessages struct {
	// The channel which the response is associated with.
	//
	// Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#channel GoogleDialogflowCxPage#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// conversation_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#conversation_success GoogleDialogflowCxPage#conversation_success}
	ConversationSuccess *GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccess `field:"optional" json:"conversationSuccess" yaml:"conversationSuccess"`
	// live_agent_handoff block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#live_agent_handoff GoogleDialogflowCxPage#live_agent_handoff}
	LiveAgentHandoff *GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoff `field:"optional" json:"liveAgentHandoff" yaml:"liveAgentHandoff"`
	// output_audio_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#output_audio_text GoogleDialogflowCxPage#output_audio_text}
	OutputAudioText *GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioText `field:"optional" json:"outputAudioText" yaml:"outputAudioText"`
	// A custom, platform-specific payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#payload GoogleDialogflowCxPage#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// play_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#play_audio GoogleDialogflowCxPage#play_audio}
	PlayAudio *GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudio `field:"optional" json:"playAudio" yaml:"playAudio"`
	// telephony_transfer_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#telephony_transfer_call GoogleDialogflowCxPage#telephony_transfer_call}
	TelephonyTransferCall *GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCall `field:"optional" json:"telephonyTransferCall" yaml:"telephonyTransferCall"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dialogflow_cx_page#text GoogleDialogflowCxPage#text}
	Text *GoogleDialogflowCxPageEntryFulfillmentMessagesText `field:"optional" json:"text" yaml:"text"`
}
