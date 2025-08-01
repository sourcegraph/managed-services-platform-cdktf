package containeranalysisnoteiammember


type ContainerAnalysisNoteIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_note_iam_member#expression ContainerAnalysisNoteIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_note_iam_member#title ContainerAnalysisNoteIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_analysis_note_iam_member#description ContainerAnalysisNoteIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

