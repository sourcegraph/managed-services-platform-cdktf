package googleserviceaccountiambinding


type GoogleServiceAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_service_account_iam_binding#expression GoogleServiceAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_service_account_iam_binding#title GoogleServiceAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_service_account_iam_binding#description GoogleServiceAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

