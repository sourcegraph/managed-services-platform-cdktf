package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig struct {
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_key GoogleDataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// Common alphabets. Only one of this, 'custom_alphabet' or 'radix' must be specified. Possible values: ["NUMERIC", "HEXADECIMAL", "UPPER_CASE_ALPHA_NUMERIC", "ALPHA_NUMERIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#common_alphabet GoogleDataLossPreventionDeidentifyTemplate#common_alphabet}
	CommonAlphabet *string `field:"optional" json:"commonAlphabet" yaml:"commonAlphabet"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#context GoogleDataLossPreventionDeidentifyTemplate#context}
	Context *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext `field:"optional" json:"context" yaml:"context"`
	// This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports.
	//
	// This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range \[2, 95\]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is:
	//
	// ''0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~'!@#$%^&*()_-+={[}]|:;"'<,>.?/''. Only one of this, 'common_alphabet' or 'radix' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#custom_alphabet GoogleDataLossPreventionDeidentifyTemplate#custom_alphabet}
	CustomAlphabet *string `field:"optional" json:"customAlphabet" yaml:"customAlphabet"`
	// The native way to select the alphabet.
	//
	// Must be in the range \[2, 95\]. Only one of this, 'custom_alphabet' or 'common_alphabet' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#radix GoogleDataLossPreventionDeidentifyTemplate#radix}
	Radix *float64 `field:"optional" json:"radix" yaml:"radix"`
	// surrogate_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#surrogate_info_type GoogleDataLossPreventionDeidentifyTemplate#surrogate_info_type}
	SurrogateInfoType *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType `field:"optional" json:"surrogateInfoType" yaml:"surrogateInfoType"`
}

