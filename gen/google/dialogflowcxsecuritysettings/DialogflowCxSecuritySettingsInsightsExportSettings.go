package dialogflowcxsecuritysettings


type DialogflowCxSecuritySettingsInsightsExportSettings struct {
	// If enabled, we will automatically exports conversations to Insights and Insights runs its analyzers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_security_settings#enable_insights_export DialogflowCxSecuritySettings#enable_insights_export}
	EnableInsightsExport interface{} `field:"required" json:"enableInsightsExport" yaml:"enableInsightsExport"`
}

