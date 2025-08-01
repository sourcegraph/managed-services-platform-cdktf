package googlebigqueryanalyticshubdataexchange

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryAnalyticsHubDataExchangeConfig struct {
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
	// The ID of the data exchange.
	//
	// Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#data_exchange_id GoogleBigqueryAnalyticsHubDataExchange#data_exchange_id}
	DataExchangeId *string `field:"required" json:"dataExchangeId" yaml:"dataExchangeId"`
	// Human-readable display name of the data exchange.
	//
	// The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#display_name GoogleBigqueryAnalyticsHubDataExchange#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this data exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#location GoogleBigqueryAnalyticsHubDataExchange#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Description of the data exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#description GoogleBigqueryAnalyticsHubDataExchange#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Documentation describing the data exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#documentation GoogleBigqueryAnalyticsHubDataExchange#documentation}
	Documentation *string `field:"optional" json:"documentation" yaml:"documentation"`
	// Base64 encoded image representing the data exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#icon GoogleBigqueryAnalyticsHubDataExchange#icon}
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#id GoogleBigqueryAnalyticsHubDataExchange#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user.
	//
	// Once enabled, this setting cannot be turned off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#log_linked_dataset_query_user_email GoogleBigqueryAnalyticsHubDataExchange#log_linked_dataset_query_user_email}
	LogLinkedDatasetQueryUserEmail interface{} `field:"optional" json:"logLinkedDatasetQueryUserEmail" yaml:"logLinkedDatasetQueryUserEmail"`
	// Email or URL of the primary point of contact of the data exchange.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#primary_contact GoogleBigqueryAnalyticsHubDataExchange#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#project GoogleBigqueryAnalyticsHubDataExchange#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// sharing_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#sharing_environment_config GoogleBigqueryAnalyticsHubDataExchange#sharing_environment_config}
	SharingEnvironmentConfig *GoogleBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig `field:"optional" json:"sharingEnvironmentConfig" yaml:"sharingEnvironmentConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_analytics_hub_data_exchange#timeouts GoogleBigqueryAnalyticsHubDataExchange#timeouts}
	Timeouts *GoogleBigqueryAnalyticsHubDataExchangeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

