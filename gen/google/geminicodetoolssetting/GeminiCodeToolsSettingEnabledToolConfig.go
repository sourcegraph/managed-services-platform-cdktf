package geminicodetoolssetting


type GeminiCodeToolsSettingEnabledToolConfig struct {
	// Key of the configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gemini_code_tools_setting#key GeminiCodeToolsSetting#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gemini_code_tools_setting#value GeminiCodeToolsSetting#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

