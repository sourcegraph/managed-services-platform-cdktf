package googledatacatalogentrygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataCatalogEntryGroupConfig struct {
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
	// The id of the entry group to create.
	//
	// The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#entry_group_id GoogleDataCatalogEntryGroup#entry_group_id}
	EntryGroupId *string `field:"required" json:"entryGroupId" yaml:"entryGroupId"`
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#description GoogleDataCatalogEntryGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#display_name GoogleDataCatalogEntryGroup#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#id GoogleDataCatalogEntryGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#project GoogleDataCatalogEntryGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// EntryGroup location region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#region GoogleDataCatalogEntryGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry_group#timeouts GoogleDataCatalogEntryGroup#timeouts}
	Timeouts *GoogleDataCatalogEntryGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

