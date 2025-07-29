package googleapihubplugininstance


type GoogleApihubPluginInstanceActionsCurationConfig struct {
	// Possible values: CURATION_TYPE_UNSPECIFIED DEFAULT_CURATION_FOR_API_METADATA CUSTOM_CURATION_FOR_API_METADATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apihub_plugin_instance#curation_type GoogleApihubPluginInstance#curation_type}
	CurationType *string `field:"optional" json:"curationType" yaml:"curationType"`
	// custom_curation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apihub_plugin_instance#custom_curation GoogleApihubPluginInstance#custom_curation}
	CustomCuration *GoogleApihubPluginInstanceActionsCurationConfigCustomCuration `field:"optional" json:"customCuration" yaml:"customCuration"`
}

