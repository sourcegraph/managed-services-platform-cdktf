package computeinstancetemplateiambinding


type ComputeInstanceTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_template_iam_binding#expression ComputeInstanceTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_template_iam_binding#title ComputeInstanceTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_template_iam_binding#description ComputeInstanceTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

