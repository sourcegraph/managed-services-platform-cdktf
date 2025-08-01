package pubsubschemaiammember


type PubsubSchemaIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/pubsub_schema_iam_member#expression PubsubSchemaIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/pubsub_schema_iam_member#title PubsubSchemaIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/pubsub_schema_iam_member#description PubsubSchemaIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

