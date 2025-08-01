package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobMappingPipelineJob struct {
	// mapping_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_pipeline_job#mapping_config GoogleHealthcarePipelineJob#mapping_config}
	MappingConfig *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig `field:"required" json:"mappingConfig" yaml:"mappingConfig"`
	// If set, the mapping pipeline will write snapshots to this FHIR store without assigning stable IDs.
	//
	// You must
	// grant your pipeline project's Cloud Healthcare Service
	// Agent serviceaccount healthcare.fhirResources.executeBundle
	// and healthcare.fhirResources.create permissions on the
	// destination store. The destination store must set
	// [disableReferentialIntegrity][FhirStore.disable_referential_integrity]
	// to true. The destination store must use FHIR version R4.
	// Format: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{fhirStoreID}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_pipeline_job#fhir_store_destination GoogleHealthcarePipelineJob#fhir_store_destination}
	FhirStoreDestination *string `field:"optional" json:"fhirStoreDestination" yaml:"fhirStoreDestination"`
	// fhir_streaming_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_pipeline_job#fhir_streaming_source GoogleHealthcarePipelineJob#fhir_streaming_source}
	FhirStreamingSource *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource `field:"optional" json:"fhirStreamingSource" yaml:"fhirStreamingSource"`
	// If set to true, a mapping pipeline will send output snapshots to the reconciliation pipeline in its dataset.
	//
	// A reconciliation
	// pipeline must exist in this dataset before a mapping pipeline
	// with a reconciliation destination can be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_pipeline_job#reconciliation_destination GoogleHealthcarePipelineJob#reconciliation_destination}
	ReconciliationDestination interface{} `field:"optional" json:"reconciliationDestination" yaml:"reconciliationDestination"`
}

