package googlecomputeinstancetemplateiammember


type GoogleComputeInstanceTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_instance_template_iam_member#expression GoogleComputeInstanceTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_instance_template_iam_member#title GoogleComputeInstanceTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_instance_template_iam_member#description GoogleComputeInstanceTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

