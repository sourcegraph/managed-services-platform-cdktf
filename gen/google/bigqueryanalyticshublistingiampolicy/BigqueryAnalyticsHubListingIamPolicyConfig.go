package bigqueryanalyticshublistingiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryAnalyticsHubListingIamPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#data_exchange_id BigqueryAnalyticsHubListingIamPolicy#data_exchange_id}.
	DataExchangeId *string `field:"required" json:"dataExchangeId" yaml:"dataExchangeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#listing_id BigqueryAnalyticsHubListingIamPolicy#listing_id}.
	ListingId *string `field:"required" json:"listingId" yaml:"listingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#policy_data BigqueryAnalyticsHubListingIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#id BigqueryAnalyticsHubListingIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#location BigqueryAnalyticsHubListingIamPolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_analytics_hub_listing_iam_policy#project BigqueryAnalyticsHubListingIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

