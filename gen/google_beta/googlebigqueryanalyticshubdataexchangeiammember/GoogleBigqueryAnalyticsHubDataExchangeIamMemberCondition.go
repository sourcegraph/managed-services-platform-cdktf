package googlebigqueryanalyticshubdataexchangeiammember


type GoogleBigqueryAnalyticsHubDataExchangeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_bigquery_analytics_hub_data_exchange_iam_member#expression GoogleBigqueryAnalyticsHubDataExchangeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_bigquery_analytics_hub_data_exchange_iam_member#title GoogleBigqueryAnalyticsHubDataExchangeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_bigquery_analytics_hub_data_exchange_iam_member#description GoogleBigqueryAnalyticsHubDataExchangeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

