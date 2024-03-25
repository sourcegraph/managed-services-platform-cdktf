package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig struct {
	// The part of the time to keep. Possible values: ["YEAR", "MONTH", "DAY_OF_MONTH", "DAY_OF_WEEK", "WEEK_OF_YEAR", "HOUR_OF_DAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_loss_prevention_deidentify_template#part_to_extract GoogleDataLossPreventionDeidentifyTemplate#part_to_extract}
	PartToExtract *string `field:"optional" json:"partToExtract" yaml:"partToExtract"`
}

