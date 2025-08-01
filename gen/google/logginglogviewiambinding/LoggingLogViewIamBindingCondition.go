package logginglogviewiambinding


type LoggingLogViewIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view_iam_binding#expression LoggingLogViewIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view_iam_binding#title LoggingLogViewIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view_iam_binding#description LoggingLogViewIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

