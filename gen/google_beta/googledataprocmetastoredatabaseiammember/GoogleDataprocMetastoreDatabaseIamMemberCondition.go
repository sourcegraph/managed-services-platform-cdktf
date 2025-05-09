package googledataprocmetastoredatabaseiammember


type GoogleDataprocMetastoreDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataproc_metastore_database_iam_member#expression GoogleDataprocMetastoreDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataproc_metastore_database_iam_member#title GoogleDataprocMetastoreDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataproc_metastore_database_iam_member#description GoogleDataprocMetastoreDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

