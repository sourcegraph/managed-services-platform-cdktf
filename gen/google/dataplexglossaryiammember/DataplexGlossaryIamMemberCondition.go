package dataplexglossaryiammember


type DataplexGlossaryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_glossary_iam_member#expression DataplexGlossaryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_glossary_iam_member#title DataplexGlossaryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_glossary_iam_member#description DataplexGlossaryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

