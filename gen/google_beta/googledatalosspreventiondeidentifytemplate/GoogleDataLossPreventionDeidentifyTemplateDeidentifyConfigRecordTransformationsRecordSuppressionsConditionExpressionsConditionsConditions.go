package googledatalosspreventiondeidentifytemplate


type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditions struct {
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_loss_prevention_deidentify_template#field GoogleDataLossPreventionDeidentifyTemplate#field}
	Field *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsField `field:"required" json:"field" yaml:"field"`
	// Operator used to compare the field or infoType to the value.
	//
	// Possible values: ["EQUAL_TO", "NOT_EQUAL_TO", "GREATER_THAN", "LESS_THAN", "GREATER_THAN_OR_EQUALS", "LESS_THAN_OR_EQUALS", "EXISTS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_loss_prevention_deidentify_template#operator GoogleDataLossPreventionDeidentifyTemplate#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_loss_prevention_deidentify_template#value GoogleDataLossPreventionDeidentifyTemplate#value}
	Value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsRecordSuppressionsConditionExpressionsConditionsConditionsValue `field:"optional" json:"value" yaml:"value"`
}

