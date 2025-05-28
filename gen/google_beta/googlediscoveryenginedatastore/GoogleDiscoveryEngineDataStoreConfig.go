package googlediscoveryenginedatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineDataStoreConfig struct {
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
	// The content config of the data store. Possible values: ["NO_CONTENT", "CONTENT_REQUIRED", "PUBLIC_WEBSITE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#content_config GoogleDiscoveryEngineDataStore#content_config}
	ContentConfig *string `field:"required" json:"contentConfig" yaml:"contentConfig"`
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#data_store_id GoogleDiscoveryEngineDataStore#data_store_id}
	DataStoreId *string `field:"required" json:"dataStoreId" yaml:"dataStoreId"`
	// The display name of the data store.
	//
	// This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#display_name GoogleDiscoveryEngineDataStore#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The industry vertical that the data store registers. Possible values: ["GENERIC", "MEDIA", "HEALTHCARE_FHIR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#industry_vertical GoogleDiscoveryEngineDataStore#industry_vertical}
	IndustryVertical *string `field:"required" json:"industryVertical" yaml:"industryVertical"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#location GoogleDiscoveryEngineDataStore#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// advanced_site_search_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#advanced_site_search_config GoogleDiscoveryEngineDataStore#advanced_site_search_config}
	AdvancedSiteSearchConfig *GoogleDiscoveryEngineDataStoreAdvancedSiteSearchConfig `field:"optional" json:"advancedSiteSearchConfig" yaml:"advancedSiteSearchConfig"`
	// If true, an advanced data store for site search will be created.
	//
	// If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#create_advanced_site_search GoogleDiscoveryEngineDataStore#create_advanced_site_search}
	CreateAdvancedSiteSearch interface{} `field:"optional" json:"createAdvancedSiteSearch" yaml:"createAdvancedSiteSearch"`
	// document_processing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#document_processing_config GoogleDiscoveryEngineDataStore#document_processing_config}
	DocumentProcessingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfig `field:"optional" json:"documentProcessingConfig" yaml:"documentProcessingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#id GoogleDiscoveryEngineDataStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#project GoogleDiscoveryEngineDataStore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A boolean flag indicating whether to skip the default schema creation for the data store.
	//
	// Only enable this flag if you are certain that the default
	// schema is incompatible with your use case.
	// If set to true, you must manually create a schema for the data store
	// before any documents can be ingested.
	// This flag cannot be specified if 'data_store.starting_schema' is
	// specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#skip_default_schema_creation GoogleDiscoveryEngineDataStore#skip_default_schema_creation}
	SkipDefaultSchemaCreation interface{} `field:"optional" json:"skipDefaultSchemaCreation" yaml:"skipDefaultSchemaCreation"`
	// The solutions that the data store enrolls. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH", "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#solution_types GoogleDiscoveryEngineDataStore#solution_types}
	SolutionTypes *[]*string `field:"optional" json:"solutionTypes" yaml:"solutionTypes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_discovery_engine_data_store#timeouts GoogleDiscoveryEngineDataStore#timeouts}
	Timeouts *GoogleDiscoveryEngineDataStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

