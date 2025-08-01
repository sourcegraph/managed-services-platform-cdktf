package googledialogflowcxpage


type GoogleDialogflowCxPageFormParameters struct {
	// advanced_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#advanced_settings GoogleDialogflowCxPage#advanced_settings}
	AdvancedSettings *GoogleDialogflowCxPageFormParametersAdvancedSettings `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// The default value of an optional parameter. If the parameter is required, the default value will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#default_value GoogleDialogflowCxPage#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The human-readable name of the parameter, unique within the form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#display_name GoogleDialogflowCxPage#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The entity type of the parameter.
	//
	// Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#entity_type GoogleDialogflowCxPage#entity_type}
	EntityType *string `field:"optional" json:"entityType" yaml:"entityType"`
	// fill_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#fill_behavior GoogleDialogflowCxPage#fill_behavior}
	FillBehavior *GoogleDialogflowCxPageFormParametersFillBehavior `field:"optional" json:"fillBehavior" yaml:"fillBehavior"`
	// Indicates whether the parameter represents a list of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#is_list GoogleDialogflowCxPage#is_list}
	IsList interface{} `field:"optional" json:"isList" yaml:"isList"`
	// Indicates whether the parameter content should be redacted in log.
	//
	// If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#redact GoogleDialogflowCxPage#redact}
	Redact interface{} `field:"optional" json:"redact" yaml:"redact"`
	// Indicates whether the parameter is required.
	//
	// Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	// Required parameters must be filled before form filling concludes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_cx_page#required GoogleDialogflowCxPage#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

