package googlelogginglogviewiambinding


type GoogleLoggingLogViewIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_logging_log_view_iam_binding#expression GoogleLoggingLogViewIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_logging_log_view_iam_binding#title GoogleLoggingLogViewIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_logging_log_view_iam_binding#description GoogleLoggingLogViewIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

