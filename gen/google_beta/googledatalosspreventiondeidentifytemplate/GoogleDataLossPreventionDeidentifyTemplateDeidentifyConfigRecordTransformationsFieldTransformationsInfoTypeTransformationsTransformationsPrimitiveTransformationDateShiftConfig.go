package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig struct {
	// For example, -5 means shift date to at most 5 days back in the past.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#lower_bound_days GoogleDataLossPreventionDeidentifyTemplate#lower_bound_days}
	LowerBoundDays *float64 `field:"required" json:"lowerBoundDays" yaml:"lowerBoundDays"`
	// Range of shift in days.
	//
	// Actual shift will be selected at random within this range (inclusive ends). Negative means shift to earlier in time. Must not be more than 365250 days (1000 years) each direction.
	//
	// For example, 3 means shift date to at most 3 days into the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#upper_bound_days GoogleDataLossPreventionDeidentifyTemplate#upper_bound_days}
	UpperBoundDays *float64 `field:"required" json:"upperBoundDays" yaml:"upperBoundDays"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#context GoogleDataLossPreventionDeidentifyTemplate#context}
	Context *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigContext `field:"optional" json:"context" yaml:"context"`
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#crypto_key GoogleDataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKey `field:"optional" json:"cryptoKey" yaml:"cryptoKey"`
}

