package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment struct {
	// advanced_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#advanced_settings GoogleDialogflowCxPage#advanced_settings}
	AdvancedSettings *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// conditional_cases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#conditional_cases GoogleDialogflowCxPage#conditional_cases}
	ConditionalCases interface{} `field:"optional" json:"conditionalCases" yaml:"conditionalCases"`
	// If the flag is true, the agent will utilize LLM to generate a text response.
	//
	// If LLM generation fails, the defined responses in the fulfillment will be respected. This flag is only useful for fulfillments associated with no-match event handlers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#enable_generative_fallback GoogleDialogflowCxPage#enable_generative_fallback}
	EnableGenerativeFallback interface{} `field:"optional" json:"enableGenerativeFallback" yaml:"enableGenerativeFallback"`
	// messages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#messages GoogleDialogflowCxPage#messages}
	Messages interface{} `field:"optional" json:"messages" yaml:"messages"`
	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs.
	//
	// If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#return_partial_responses GoogleDialogflowCxPage#return_partial_responses}
	ReturnPartialResponses interface{} `field:"optional" json:"returnPartialResponses" yaml:"returnPartialResponses"`
	// set_parameter_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#set_parameter_actions GoogleDialogflowCxPage#set_parameter_actions}
	SetParameterActions interface{} `field:"optional" json:"setParameterActions" yaml:"setParameterActions"`
	// The tag used by the webhook to identify which fulfillment is being called.
	//
	// This field is required if webhook is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#tag GoogleDialogflowCxPage#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dialogflow_cx_page#webhook GoogleDialogflowCxPage#webhook}
	Webhook *string `field:"optional" json:"webhook" yaml:"webhook"`
}

