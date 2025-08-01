package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingBigqueryDatasetSelectedResources struct {
	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing#table BigqueryAnalyticsHubListing#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
}

