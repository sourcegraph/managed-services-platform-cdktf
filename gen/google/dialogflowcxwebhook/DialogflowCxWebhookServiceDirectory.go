package dialogflowcxwebhook


type DialogflowCxWebhookServiceDirectory struct {
	// The name of Service Directory service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_webhook#service DialogflowCxWebhook#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_webhook#generic_web_service DialogflowCxWebhook#generic_web_service}
	GenericWebService *DialogflowCxWebhookServiceDirectoryGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
}

