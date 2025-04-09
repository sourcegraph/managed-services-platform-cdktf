package dataprocmetastoretableiammember


type DataprocMetastoreTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_metastore_table_iam_member#expression DataprocMetastoreTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_metastore_table_iam_member#title DataprocMetastoreTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_metastore_table_iam_member#description DataprocMetastoreTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

