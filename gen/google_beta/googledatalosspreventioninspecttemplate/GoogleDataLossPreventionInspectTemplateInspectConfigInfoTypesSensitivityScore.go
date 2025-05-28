package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigInfoTypesSensitivityScore struct {
	// The sensitivity score applied to the resource. Possible values: ["SENSITIVITY_LOW", "SENSITIVITY_MODERATE", "SENSITIVITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_data_loss_prevention_inspect_template#score GoogleDataLossPreventionInspectTemplate#score}
	Score *string `field:"required" json:"score" yaml:"score"`
}

