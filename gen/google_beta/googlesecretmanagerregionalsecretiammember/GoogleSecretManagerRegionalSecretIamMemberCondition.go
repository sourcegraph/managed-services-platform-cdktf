package googlesecretmanagerregionalsecretiammember


type GoogleSecretManagerRegionalSecretIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_secret_manager_regional_secret_iam_member#expression GoogleSecretManagerRegionalSecretIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_secret_manager_regional_secret_iam_member#title GoogleSecretManagerRegionalSecretIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_secret_manager_regional_secret_iam_member#description GoogleSecretManagerRegionalSecretIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

