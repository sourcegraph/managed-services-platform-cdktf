package googlesecretmanagersecretiambinding


type GoogleSecretManagerSecretIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_secret_manager_secret_iam_binding#expression GoogleSecretManagerSecretIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_secret_manager_secret_iam_binding#title GoogleSecretManagerSecretIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_secret_manager_secret_iam_binding#description GoogleSecretManagerSecretIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

