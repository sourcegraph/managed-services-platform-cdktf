package datacloudflareorganizations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareOrganizationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#containing DataCloudflareOrganizations#containing}.
	Containing *DataCloudflareOrganizationsContaining `field:"optional" json:"containing" yaml:"containing"`
	// Only return organizations with the specified IDs (ex. id=foo&id=bar). Send multiple elements by repeating the query value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#id DataCloudflareOrganizations#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#max_items DataCloudflareOrganizations#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#name DataCloudflareOrganizations#name}.
	Name *DataCloudflareOrganizationsName `field:"optional" json:"name" yaml:"name"`
	// The amount of items to return. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#page_size DataCloudflareOrganizations#page_size}
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// An opaque token returned from the last list response that when provided will retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent
	// requests with a page token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#page_token DataCloudflareOrganizations#page_token}
	PageToken *string `field:"optional" json:"pageToken" yaml:"pageToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#parent DataCloudflareOrganizations#parent}.
	Parent *DataCloudflareOrganizationsParent `field:"optional" json:"parent" yaml:"parent"`
}

