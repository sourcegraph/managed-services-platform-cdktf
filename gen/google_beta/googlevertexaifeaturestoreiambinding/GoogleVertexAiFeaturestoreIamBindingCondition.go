package googlevertexaifeaturestoreiambinding


type GoogleVertexAiFeaturestoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vertex_ai_featurestore_iam_binding#expression GoogleVertexAiFeaturestoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vertex_ai_featurestore_iam_binding#title GoogleVertexAiFeaturestoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vertex_ai_featurestore_iam_binding#description GoogleVertexAiFeaturestoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

