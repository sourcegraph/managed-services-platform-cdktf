package dialogflowcxflow


type DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText struct {
	// The SSML text to be synthesized. For more information, see SSML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_flow#ssml DialogflowCxFlow#ssml}
	Ssml *string `field:"optional" json:"ssml" yaml:"ssml"`
	// The raw text to be synthesized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_flow#text DialogflowCxFlow#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

