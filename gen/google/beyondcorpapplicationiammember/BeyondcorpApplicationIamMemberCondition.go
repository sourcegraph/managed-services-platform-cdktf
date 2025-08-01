package beyondcorpapplicationiammember


type BeyondcorpApplicationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_application_iam_member#expression BeyondcorpApplicationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_application_iam_member#title BeyondcorpApplicationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_application_iam_member#description BeyondcorpApplicationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

