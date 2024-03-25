package googlebigquerydatasetiambinding


type GoogleBigqueryDatasetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_bigquery_dataset_iam_binding#expression GoogleBigqueryDatasetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_bigquery_dataset_iam_binding#title GoogleBigqueryDatasetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_bigquery_dataset_iam_binding#description GoogleBigqueryDatasetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

