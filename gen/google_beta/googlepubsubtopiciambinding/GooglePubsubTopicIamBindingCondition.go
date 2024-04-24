package googlepubsubtopiciambinding


type GooglePubsubTopicIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_pubsub_topic_iam_binding#expression GooglePubsubTopicIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_pubsub_topic_iam_binding#title GooglePubsubTopicIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_pubsub_topic_iam_binding#description GooglePubsubTopicIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

