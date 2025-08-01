package gkehubscopeiambinding


type GkeHubScopeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope_iam_binding#expression GkeHubScopeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope_iam_binding#title GkeHubScopeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_scope_iam_binding#description GkeHubScopeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

