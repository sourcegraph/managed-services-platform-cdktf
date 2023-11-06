package googletagstagkeyiambinding


type GoogleTagsTagKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_tags_tag_key_iam_binding#expression GoogleTagsTagKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_tags_tag_key_iam_binding#title GoogleTagsTagKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_tags_tag_key_iam_binding#description GoogleTagsTagKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

