package bigqueryconnectioniammember


type BigqueryConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_connection_iam_member#expression BigqueryConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_connection_iam_member#title BigqueryConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_connection_iam_member#description BigqueryConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

