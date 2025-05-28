package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentConditionalCases struct {
	// A JSON encoded list of cascading if-else conditions.
	//
	// Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.
	// See [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dialogflow_cx_page#cases GoogleDialogflowCxPage#cases}
	Cases *string `field:"optional" json:"cases" yaml:"cases"`
}

