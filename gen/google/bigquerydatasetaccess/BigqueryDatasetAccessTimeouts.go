package bigquerydatasetaccess


type BigqueryDatasetAccessTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_dataset_access#create BigqueryDatasetAccessA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_dataset_access#delete BigqueryDatasetAccessA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

