package googlediscoveryenginesitemap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineSitemapConfig struct {
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
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#data_store_id GoogleDiscoveryEngineSitemap#data_store_id}
	DataStoreId *string `field:"required" json:"dataStoreId" yaml:"dataStoreId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#location GoogleDiscoveryEngineSitemap#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#id GoogleDiscoveryEngineSitemap#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#project GoogleDiscoveryEngineSitemap#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#timeouts GoogleDiscoveryEngineSitemap#timeouts}
	Timeouts *GoogleDiscoveryEngineSitemapTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Public URI for the sitemap, e.g. "www.example.com/sitemap.xml".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_sitemap#uri GoogleDiscoveryEngineSitemap#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

