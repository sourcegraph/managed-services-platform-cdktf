package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType struct {
	// Name of the information type.
	//
	// Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern '[A-Za-z0-9$-_]{1,64}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#name GoogleDataLossPreventionDeidentifyTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#sensitivity_score GoogleDataLossPreventionDeidentifyTemplate#sensitivity_score}
	SensitivityScore *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoTypeSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Optional version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_deidentify_template#version GoogleDataLossPreventionDeidentifyTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

