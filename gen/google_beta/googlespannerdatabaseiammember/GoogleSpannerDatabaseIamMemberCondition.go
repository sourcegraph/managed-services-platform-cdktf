package googlespannerdatabaseiammember


type GoogleSpannerDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_spanner_database_iam_member#expression GoogleSpannerDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_spanner_database_iam_member#title GoogleSpannerDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_spanner_database_iam_member#description GoogleSpannerDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

