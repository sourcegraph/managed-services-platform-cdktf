package dialogflowcxpage


type DialogflowCxPageFormParametersFillBehavior struct {
	// initial_prompt_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#initial_prompt_fulfillment DialogflowCxPage#initial_prompt_fulfillment}
	InitialPromptFulfillment *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment `field:"optional" json:"initialPromptFulfillment" yaml:"initialPromptFulfillment"`
	// reprompt_event_handlers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#reprompt_event_handlers DialogflowCxPage#reprompt_event_handlers}
	RepromptEventHandlers interface{} `field:"optional" json:"repromptEventHandlers" yaml:"repromptEventHandlers"`
}

