package googlehealthcarefhirstore


type GoogleHealthcareFhirStoreStreamConfigsBigqueryDestination struct {
	// BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_fhir_store#dataset_uri GoogleHealthcareFhirStore#dataset_uri}
	DatasetUri *string `field:"required" json:"datasetUri" yaml:"datasetUri"`
	// schema_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_healthcare_fhir_store#schema_config GoogleHealthcareFhirStore#schema_config}
	SchemaConfig *GoogleHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig `field:"required" json:"schemaConfig" yaml:"schemaConfig"`
}

