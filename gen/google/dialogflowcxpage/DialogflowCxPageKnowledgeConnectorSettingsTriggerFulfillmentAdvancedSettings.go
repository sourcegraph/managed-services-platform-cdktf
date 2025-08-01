package dialogflowcxpage


type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings struct {
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#dtmf_settings DialogflowCxPage#dtmf_settings}
	DtmfSettings *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
	// logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#logging_settings DialogflowCxPage#logging_settings}
	LoggingSettings *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings `field:"optional" json:"loggingSettings" yaml:"loggingSettings"`
	// speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#speech_settings DialogflowCxPage#speech_settings}
	SpeechSettings *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings `field:"optional" json:"speechSettings" yaml:"speechSettings"`
}

