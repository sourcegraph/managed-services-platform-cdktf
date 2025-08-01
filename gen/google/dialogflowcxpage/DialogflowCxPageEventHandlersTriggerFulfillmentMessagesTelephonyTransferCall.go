package dialogflowcxpage


type DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall struct {
	// Transfer the call to a phone number in E.164 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_page#phone_number DialogflowCxPage#phone_number}
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

