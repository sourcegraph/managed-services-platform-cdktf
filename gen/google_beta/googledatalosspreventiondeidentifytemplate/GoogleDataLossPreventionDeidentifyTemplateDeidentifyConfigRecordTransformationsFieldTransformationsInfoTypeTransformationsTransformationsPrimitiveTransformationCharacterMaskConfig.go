package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig struct {
	// characters_to_ignore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#characters_to_ignore GoogleDataLossPreventionDeidentifyTemplate#characters_to_ignore}
	CharactersToIgnore interface{} `field:"optional" json:"charactersToIgnore" yaml:"charactersToIgnore"`
	// Character to use to mask the sensitive values—for example, * for an alphabetic string such as a name, or 0 for a numeric string such as ZIP code or credit card number.
	//
	// This string must have a length of 1. If not supplied, this value defaults to * for
	// strings, and 0 for digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#masking_character GoogleDataLossPreventionDeidentifyTemplate#masking_character}
	MaskingCharacter *string `field:"optional" json:"maskingCharacter" yaml:"maskingCharacter"`
	// Number of characters to mask.
	//
	// If not set, all matching chars will be masked. Skipped characters do not count towards this tally.
	// If number_to_mask is negative, this denotes inverse masking. Cloud DLP masks all but a number of characters. For example, suppose you have the following values:
	// - 'masking_character' is *
	// - 'number_to_mask' is -4
	// - 'reverse_order' is false
	// - 'characters_to_ignore' includes -
	// - Input string is 1234-5678-9012-3456
	//
	// The resulting de-identified string is ****-****-****-3456. Cloud DLP masks all but the last four characters. If reverseOrder is true, all but the first four characters are masked as 1234-****-****-****.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#number_to_mask GoogleDataLossPreventionDeidentifyTemplate#number_to_mask}
	NumberToMask *float64 `field:"optional" json:"numberToMask" yaml:"numberToMask"`
	// Mask characters in reverse order.
	//
	// For example, if masking_character is 0, number_to_mask is 14, and reverse_order is 'false', then the
	// input string '1234-5678-9012-3456' is masked as '00000000000000-3456'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#reverse_order GoogleDataLossPreventionDeidentifyTemplate#reverse_order}
	ReverseOrder interface{} `field:"optional" json:"reverseOrder" yaml:"reverseOrder"`
}

