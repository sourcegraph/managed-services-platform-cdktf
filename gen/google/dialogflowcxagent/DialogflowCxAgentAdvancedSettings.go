package dialogflowcxagent


type DialogflowCxAgentAdvancedSettings struct {
	// audio_export_gcs_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_agent#audio_export_gcs_destination DialogflowCxAgent#audio_export_gcs_destination}
	AudioExportGcsDestination *DialogflowCxAgentAdvancedSettingsAudioExportGcsDestination `field:"optional" json:"audioExportGcsDestination" yaml:"audioExportGcsDestination"`
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_agent#dtmf_settings DialogflowCxAgent#dtmf_settings}
	DtmfSettings *DialogflowCxAgentAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
	// logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_agent#logging_settings DialogflowCxAgent#logging_settings}
	LoggingSettings *DialogflowCxAgentAdvancedSettingsLoggingSettings `field:"optional" json:"loggingSettings" yaml:"loggingSettings"`
	// speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_agent#speech_settings DialogflowCxAgent#speech_settings}
	SpeechSettings *DialogflowCxAgentAdvancedSettingsSpeechSettings `field:"optional" json:"speechSettings" yaml:"speechSettings"`
}

