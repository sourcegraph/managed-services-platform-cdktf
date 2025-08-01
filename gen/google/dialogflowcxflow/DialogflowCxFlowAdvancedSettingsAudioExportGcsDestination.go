package dialogflowcxflow


type DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination struct {
	// The Google Cloud Storage URI for the exported objects.
	//
	// Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	// Format: gs://bucket/object-name-or-prefix
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_flow#uri DialogflowCxFlow#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

