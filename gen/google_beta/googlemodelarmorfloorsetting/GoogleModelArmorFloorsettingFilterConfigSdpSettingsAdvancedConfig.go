package googlemodelarmorfloorsetting


type GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig struct {
	// Optional Sensitive Data Protection Deidentify template resource name.
	//
	// If provided then DeidentifyContent action is performed during Sanitization
	// using this template and inspect template. The De-identified data will
	// be returned in SdpDeidentifyResult.
	// Note that all info-types present in the deidentify template must be present
	// in inspect template.
	//
	// e.g.
	// 'projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_model_armor_floorsetting#deidentify_template GoogleModelArmorFloorsetting#deidentify_template}
	DeidentifyTemplate *string `field:"optional" json:"deidentifyTemplate" yaml:"deidentifyTemplate"`
	// Sensitive Data Protection inspect template resource name.
	//
	// If only inspect template is provided (de-identify template not provided),
	// then Sensitive Data Protection InspectContent action is performed during
	// Sanitization. All Sensitive Data Protection findings identified during
	// inspection will be returned as SdpFinding in SdpInsepctionResult.
	//
	// e.g:-
	// 'projects/{project}/locations/{location}/inspectTemplates/{inspect_template}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_model_armor_floorsetting#inspect_template GoogleModelArmorFloorsetting#inspect_template}
	InspectTemplate *string `field:"optional" json:"inspectTemplate" yaml:"inspectTemplate"`
}

