package googlepubsubschemaiammember


type GooglePubsubSchemaIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_schema_iam_member#expression GooglePubsubSchemaIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_schema_iam_member#title GooglePubsubSchemaIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_schema_iam_member#description GooglePubsubSchemaIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

