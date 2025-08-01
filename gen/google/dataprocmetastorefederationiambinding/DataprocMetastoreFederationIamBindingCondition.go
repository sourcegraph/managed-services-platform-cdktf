package dataprocmetastorefederationiambinding


type DataprocMetastoreFederationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_metastore_federation_iam_binding#expression DataprocMetastoreFederationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_metastore_federation_iam_binding#title DataprocMetastoreFederationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_metastore_federation_iam_binding#description DataprocMetastoreFederationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

