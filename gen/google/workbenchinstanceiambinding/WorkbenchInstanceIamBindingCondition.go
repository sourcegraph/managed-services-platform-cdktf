package workbenchinstanceiambinding


type WorkbenchInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workbench_instance_iam_binding#expression WorkbenchInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workbench_instance_iam_binding#title WorkbenchInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workbench_instance_iam_binding#description WorkbenchInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

