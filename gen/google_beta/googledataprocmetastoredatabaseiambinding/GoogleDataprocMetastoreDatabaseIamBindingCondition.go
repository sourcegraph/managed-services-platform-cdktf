package googledataprocmetastoredatabaseiambinding


type GoogleDataprocMetastoreDatabaseIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_dataproc_metastore_database_iam_binding#expression GoogleDataprocMetastoreDatabaseIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_dataproc_metastore_database_iam_binding#title GoogleDataprocMetastoreDatabaseIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_dataproc_metastore_database_iam_binding#description GoogleDataprocMetastoreDatabaseIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

