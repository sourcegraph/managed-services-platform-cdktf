package googledataplexentrytypeiambinding


type GoogleDataplexEntryTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataplex_entry_type_iam_binding#expression GoogleDataplexEntryTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataplex_entry_type_iam_binding#title GoogleDataplexEntryTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_dataplex_entry_type_iam_binding#description GoogleDataplexEntryTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

