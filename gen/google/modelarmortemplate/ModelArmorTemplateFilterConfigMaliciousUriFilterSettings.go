package modelarmortemplate


type ModelArmorTemplateFilterConfigMaliciousUriFilterSettings struct {
	// Tells whether the Malicious URI filter is enabled or disabled. Possible values: ENABLED DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/model_armor_template#filter_enforcement ModelArmorTemplate#filter_enforcement}
	FilterEnforcement *string `field:"optional" json:"filterEnforcement" yaml:"filterEnforcement"`
}

