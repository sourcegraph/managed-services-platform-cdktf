package googledatacatalogtaxonomy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataCatalogTaxonomyConfig struct {
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
	// User defined name of this taxonomy.
	//
	// The taxonomy display name must be unique within an organization.
	// It must: contain only unicode letters, numbers, underscores, dashes
	// and spaces; not start or end with spaces; and be at most 200 bytes
	// long when encoded in UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#display_name GoogleDataCatalogTaxonomy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// A list of policy types that are activated for this taxonomy.
	//
	// If not set,
	// defaults to an empty list. Possible values: ["POLICY_TYPE_UNSPECIFIED", "FINE_GRAINED_ACCESS_CONTROL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#activated_policy_types GoogleDataCatalogTaxonomy#activated_policy_types}
	ActivatedPolicyTypes *[]*string `field:"optional" json:"activatedPolicyTypes" yaml:"activatedPolicyTypes"`
	// Description of this taxonomy.
	//
	// It must: contain only unicode characters,
	// tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
	// long when encoded in UTF-8. If not set, defaults to an empty description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#description GoogleDataCatalogTaxonomy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#id GoogleDataCatalogTaxonomy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#project GoogleDataCatalogTaxonomy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Taxonomy location region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#region GoogleDataCatalogTaxonomy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_taxonomy#timeouts GoogleDataCatalogTaxonomy#timeouts}
	Timeouts *GoogleDataCatalogTaxonomyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

