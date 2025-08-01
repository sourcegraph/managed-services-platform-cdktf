package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfigCharactersToIgnore struct {
	// Characters to not transform when masking. Only one of this or 'common_characters_to_ignore' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#characters_to_skip GoogleDataLossPreventionDeidentifyTemplate#characters_to_skip}
	CharactersToSkip *string `field:"optional" json:"charactersToSkip" yaml:"charactersToSkip"`
	// Common characters to not transform when masking.
	//
	// Useful to avoid removing punctuation. Only one of this or 'characters_to_skip' must be specified. Possible values: ["NUMERIC", "ALPHA_UPPER_CASE", "ALPHA_LOWER_CASE", "PUNCTUATION", "WHITESPACE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#common_characters_to_ignore GoogleDataLossPreventionDeidentifyTemplate#common_characters_to_ignore}
	CommonCharactersToIgnore *string `field:"optional" json:"commonCharactersToIgnore" yaml:"commonCharactersToIgnore"`
}

