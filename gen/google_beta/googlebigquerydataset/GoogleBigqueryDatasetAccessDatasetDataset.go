package googlebigquerydataset


type GoogleBigqueryDatasetAccessDatasetDataset struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_dataset#dataset_id GoogleBigqueryDataset#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_dataset#project_id GoogleBigqueryDataset#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

