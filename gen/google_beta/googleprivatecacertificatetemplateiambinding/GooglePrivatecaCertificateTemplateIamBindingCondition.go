package googleprivatecacertificatetemplateiambinding


type GooglePrivatecaCertificateTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_privateca_certificate_template_iam_binding#expression GooglePrivatecaCertificateTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_privateca_certificate_template_iam_binding#title GooglePrivatecaCertificateTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_privateca_certificate_template_iam_binding#description GooglePrivatecaCertificateTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

