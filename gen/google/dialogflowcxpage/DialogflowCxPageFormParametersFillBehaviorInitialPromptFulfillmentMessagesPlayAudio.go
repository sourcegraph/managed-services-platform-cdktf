package dialogflowcxpage


type DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudio struct {
	// URI of the audio clip.
	//
	// Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#audio_uri DialogflowCxPage#audio_uri}
	AudioUri *string `field:"required" json:"audioUri" yaml:"audioUri"`
}

