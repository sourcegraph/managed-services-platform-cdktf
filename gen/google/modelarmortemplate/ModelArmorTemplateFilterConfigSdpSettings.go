package modelarmortemplate


type ModelArmorTemplateFilterConfigSdpSettings struct {
	// advanced_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/model_armor_template#advanced_config ModelArmorTemplate#advanced_config}
	AdvancedConfig *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig `field:"optional" json:"advancedConfig" yaml:"advancedConfig"`
	// basic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/model_armor_template#basic_config ModelArmorTemplate#basic_config}
	BasicConfig *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig `field:"optional" json:"basicConfig" yaml:"basicConfig"`
}

