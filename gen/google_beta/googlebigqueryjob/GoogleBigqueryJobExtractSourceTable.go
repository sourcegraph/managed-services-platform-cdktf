package googlebigqueryjob


type GoogleBigqueryJobExtractSourceTable struct {
	// The table. Can be specified '{{table_id}}' if 'project_id' and 'dataset_id' are also set, or of the form 'projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}' if not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#table_id GoogleBigqueryJob#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#dataset_id GoogleBigqueryJob#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#project_id GoogleBigqueryJob#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

