package dialogflowcxpage


type DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActions struct {
	// Display name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#parameter DialogflowCxPage#parameter}
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The new JSON-encoded value of the parameter. A null value clears the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#value DialogflowCxPage#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

