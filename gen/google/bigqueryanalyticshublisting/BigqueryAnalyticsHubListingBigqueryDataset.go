package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingBigqueryDataset struct {
	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing#dataset BigqueryAnalyticsHubListing#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// selected_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing#selected_resources BigqueryAnalyticsHubListing#selected_resources}
	SelectedResources interface{} `field:"optional" json:"selectedResources" yaml:"selectedResources"`
}

