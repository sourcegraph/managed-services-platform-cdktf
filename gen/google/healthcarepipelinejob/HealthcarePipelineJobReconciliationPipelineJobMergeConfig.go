package healthcarepipelinejob


type HealthcarePipelineJobReconciliationPipelineJobMergeConfig struct {
	// whistle_config_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/healthcare_pipeline_job#whistle_config_source HealthcarePipelineJob#whistle_config_source}
	WhistleConfigSource *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource `field:"required" json:"whistleConfigSource" yaml:"whistleConfigSource"`
	// Describes the mapping configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/healthcare_pipeline_job#description HealthcarePipelineJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

