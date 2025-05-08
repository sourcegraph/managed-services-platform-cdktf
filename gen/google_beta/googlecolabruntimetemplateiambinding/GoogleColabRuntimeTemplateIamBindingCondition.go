package googlecolabruntimetemplateiambinding


type GoogleColabRuntimeTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_colab_runtime_template_iam_binding#expression GoogleColabRuntimeTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_colab_runtime_template_iam_binding#title GoogleColabRuntimeTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_colab_runtime_template_iam_binding#description GoogleColabRuntimeTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

