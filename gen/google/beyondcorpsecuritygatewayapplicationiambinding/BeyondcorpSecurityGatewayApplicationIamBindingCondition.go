package beyondcorpsecuritygatewayapplicationiambinding


type BeyondcorpSecurityGatewayApplicationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application_iam_binding#expression BeyondcorpSecurityGatewayApplicationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application_iam_binding#title BeyondcorpSecurityGatewayApplicationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_application_iam_binding#description BeyondcorpSecurityGatewayApplicationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

