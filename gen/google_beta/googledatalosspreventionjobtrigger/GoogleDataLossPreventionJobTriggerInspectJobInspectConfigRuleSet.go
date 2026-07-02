package googledatalosspreventionjobtrigger

type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigRuleSet struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#rules GoogleDataLossPreventionJobTrigger#rules}
	Rules any `field:"required" json:"rules" yaml:"rules"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#info_types GoogleDataLossPreventionJobTrigger#info_types}
	InfoTypes any `field:"optional" json:"infoTypes" yaml:"infoTypes"`
}
