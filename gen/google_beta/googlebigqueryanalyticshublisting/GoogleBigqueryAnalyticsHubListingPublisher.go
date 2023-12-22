package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingPublisher struct {
	// Name of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_bigquery_analytics_hub_listing#name GoogleBigqueryAnalyticsHubListing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email or URL of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_bigquery_analytics_hub_listing#primary_contact GoogleBigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
}

