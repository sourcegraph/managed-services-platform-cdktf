package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped struct {
	// A 128/192/256 bit key.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_data_loss_prevention_deidentify_template#key GoogleDataLossPreventionDeidentifyTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

