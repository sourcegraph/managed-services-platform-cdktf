package googledialogflowcxflow


type GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio struct {
	// URI of the audio clip.
	//
	// Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dialogflow_cx_flow#audio_uri GoogleDialogflowCxFlow#audio_uri}
	AudioUri *string `field:"required" json:"audioUri" yaml:"audioUri"`
}

