package securesourcemanagerrepositoryiambinding


type SecureSourceManagerRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_binding#expression SecureSourceManagerRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_binding#title SecureSourceManagerRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_binding#description SecureSourceManagerRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

