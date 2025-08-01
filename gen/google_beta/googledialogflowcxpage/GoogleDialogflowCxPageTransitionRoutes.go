package googledialogflowcxpage


type GoogleDialogflowCxPageTransitionRoutes struct {
	// The condition to evaluate against form parameters or session parameters.
	//
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#condition GoogleDialogflowCxPage#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The unique identifier of an Intent.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#intent GoogleDialogflowCxPage#intent}
	Intent *string `field:"optional" json:"intent" yaml:"intent"`
	// The target flow to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#target_flow GoogleDialogflowCxPage#target_flow}
	TargetFlow *string `field:"optional" json:"targetFlow" yaml:"targetFlow"`
	// The target page to transition to. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#target_page GoogleDialogflowCxPage#target_page}
	TargetPage *string `field:"optional" json:"targetPage" yaml:"targetPage"`
	// trigger_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#trigger_fulfillment GoogleDialogflowCxPage#trigger_fulfillment}
	TriggerFulfillment *GoogleDialogflowCxPageTransitionRoutesTriggerFulfillment `field:"optional" json:"triggerFulfillment" yaml:"triggerFulfillment"`
}

