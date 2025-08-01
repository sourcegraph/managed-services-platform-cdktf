package clouddeploycustomtargettypeiammember


type ClouddeployCustomTargetTypeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_custom_target_type_iam_member#expression ClouddeployCustomTargetTypeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_custom_target_type_iam_member#title ClouddeployCustomTargetTypeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_custom_target_type_iam_member#description ClouddeployCustomTargetTypeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

