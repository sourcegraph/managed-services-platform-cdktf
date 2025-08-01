package apigeeenvironmentiambinding


type ApigeeEnvironmentIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apigee_environment_iam_binding#expression ApigeeEnvironmentIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apigee_environment_iam_binding#title ApigeeEnvironmentIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apigee_environment_iam_binding#description ApigeeEnvironmentIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

