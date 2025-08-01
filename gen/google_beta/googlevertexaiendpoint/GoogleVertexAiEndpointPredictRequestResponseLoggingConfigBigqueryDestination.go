package googlevertexaiendpoint


type GoogleVertexAiEndpointPredictRequestResponseLoggingConfigBigqueryDestination struct {
	// BigQuery URI to a project or table, up to 2000 characters long.
	//
	// When only the project is specified, the Dataset and Table is created. When the full table reference is specified, the Dataset must exist and table must not exist. Accepted forms: - BigQuery path. For example: 'bq://projectId' or 'bq://projectId.bqDatasetId' or 'bq://projectId.bqDatasetId.bqTableId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_endpoint#output_uri GoogleVertexAiEndpoint#output_uri}
	OutputUri *string `field:"optional" json:"outputUri" yaml:"outputUri"`
}

